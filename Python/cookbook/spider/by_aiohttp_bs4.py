"""
封装下使其易用

class selector:
    target_item = TextField(css_select="tr.athing")
    title = TextField(css_select="a.storylink")

spider(url).select(selector) -> {"target_item": ..., "title": ...}
"""

import asyncio
import typing as T
import aiohttp
from collections import UserDict
from bs4 import BeautifulSoup
import weakref
import functools
from lxml import etree
from asyncref.adapter import wrap


class TextField(T.NamedTuple):
    css_select: T.Optional[str] = None
    xpath_select: T.Optional[str] = None


class Selector(type):
    """
    需要记录子类所包含的TextField，这样spider才能用来解析
    """

    _fields_name: T.List[str]

    class _member(UserDict):
        def __init__(self):
            super().__init__()
            self.names = []

        def __setitem__(self, key, value):
            if isinstance(value, TextField):
                self.names.append(key)

            super().__setitem__(key, value)

    @classmethod
    def __prepare__(metacls, cls_name: str, bases: T.Tuple):
        cls_dict = metacls._member()
        return cls_dict

    def __new__(metacls, cls_name: str, bases: T.Tuple, cls_dict: _member):
        cls = type.__new__(metacls, cls_name, bases, dict(cls_dict))
        cls._fields_name = cls_dict.names
        return cls

    @property
    def fields(self) -> T.Dict[str, TextField]:
        names = self._fields_name
        return {name: getattr(self, name) for name in names}


class Spider:
    def __init__(self, url: str):
        self._url = url
        self._soup = BeautifulSoup(self.html, "html.parser")
        self._selector = etree.HTML(self.html)

    @functools.cached_property
    def html(self) -> BeautifulSoup:
        async def html_():
            async with aiohttp.ClientSession() as session:
                async with session.get(self._url, timeout=30) as response:
                    assert response.status == 200
                    return await response.read()

        return wrap(html_)

    async def select(self, selector: Selector):
        """
        基于目标url，通过selector中的textfield来解析出结果，并且返回结果
        """

        res = {}
        for name, path in selector.fields.items():
            content = None
            if path.xpath_select is None:
                content = self._selector.xpath(path.xpath_select)

            res[name] = content
        return res


if __name__ == "__main__":
    print("[Test] Selector")

    class item(metaclass=Selector):
        target_item = TextField(css_select="tr.athing")
        title = TextField(css_select="a.storylink")

    print(item.fields)

    print("[Test] Spider")
    spider = Spider("https://www.baidu.com")
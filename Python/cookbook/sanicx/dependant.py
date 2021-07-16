"""
依赖注入
"""
__all__ = ("Container", "ContainerManager", "Con", "Singleton", "Factory", "Attr")

import functools
import importlib
import typing as T

from dependency_injector import providers, containers
from dependency_injector.wiring import Provide, ProvidersMap

# alias
Container = containers.DeclarativeContainer

Con = providers.Container
Factory = providers.Factory


class SubContainer:
    def __init__(self, name: str, sub: Container):
        self.name = name
        self.sub = sub

    @classmethod
    def builder(cls, name: str):
        def _builder(sub: Container):
            ins = cls.__new__(cls)
            ins.__init__(name, sub)
            return ins

        return _builder

    def __get__(self, instance, owner):
        print(instance, owner, "get")

    def __set__(self, instance, value):
        print("set")

    def __delete__(self, instance):
        print("delete")


def Attr(module_name: str = None, name: str = None):
    """
    Return: 直接返回某个属性
    """

    def fn():
        module = importlib.import_module(module_name)
        if not hasattr(module, name):
            raise Exception()
        return getattr(module, name)

    return providers.Singleton(fn)


def Singleton(
        module_name: str = None,
        name: str = None,
        /,
        fn: T.Callable = None,
        args: T.Tuple = None,
        kwargs: T.Dict = None,
):
    """
    Args:
        module_name: 模块位置
        name: 变量名字
        fn: 如果指定了fn且为callable，那么就直接使用Singleton来生成
    Return: 返回的是某个属性实例
    """
    args = args or tuple()
    kwargs = kwargs or dict()

    if fn is None:

        def fn(*args_, **kwargs_):
            module = importlib.import_module(module_name)
            if not hasattr(module, name):
                raise Exception()
            return getattr(module, name)(*args_, **kwargs_)

    return providers.Singleton(fn, *args, **kwargs)


class ContainerManager:
    """
    依赖注入的容器管理类类
    """

    def __init__(
            self,
            container: T.Type[Container],
    ):
        self._container = container
        self._providers_map = ProvidersMap(container())  # 依赖对象的解析

    @classmethod
    def builder(cls, app: "sanix.app.SanicX", container: T.Type[Container] = None):
        """
        为container容器再添加一个获取app的对象
        """
        if container is None:
            container = containers.DeclarativeContainer
        setattr(container, "app", providers.Singleton(lambda: app))
        setattr(container, "config", providers.Singleton(lambda: app.config))

        ins = super().__new__(cls)
        ins.__init__(container)
        return ins

    def get_bean(self, marker_str: str):
        marker = Provide[
            functools.reduce(
                lambda conf, fi: getattr(conf, fi),
                marker_str.split("."),
                self._container,
            )
        ]
        provider = self._providers_map.resolve_provider(marker.provider)
        if provider is None:
            return None
        return provider()


if __name__ == '__main__':
    class A:
        @SubContainer.builder("name")
        class B:
            pass


    print(A.B)
    A().B = 2

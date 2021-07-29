"""
在内置的导入功能基础上新增通过url导入

$ python3 -m http.server 12800
Serving HTTP on 0.0.0.0 port 12800 (http://0.0.0.0:12800/) ...
"""
from importlib.abc import MetaPathFinder
from importlib.machinery import ModuleSpec
from importlib import abc
import sys
import urllib.request as urllib2


class UrlMetaLoader(abc.SourceLoader):

    def __init__(self, baseurl):
        self.baseurl = baseurl

    def get_code(self, fullname):
        f = urllib2.urlopen(self.get_filename(fullname))
        return f.read()
    
    def get_filename(self, fullname):
        return self.baseurl + fullname + '.py'

    def load_module(self, fullname):
        """
        需要你在查找器里手动执行，才能实现模块的加载
        """
        code = self.get_code(fullname)
        mod = sys.modules.setdefault(fullname, imp.new_module(fullname))
        mod.__file__ = self.get_filename(fullname)
        mod.__loader__ = self
        mod.__package__ = fullname
        exec(code, mod.__dict__)
        return None

    def get_data(self):
        pass

    def execute_module(self, module):
        """
        必须重载，而且不应该有任何逻辑，即使它并不是抽象方法
        """
        pass


class UrlMetaFinder(MetaPathFinder):

    def __init__(self, baseurl):
        self._baseurl = baseurl

    def find_spec(self, fullname, path=None, target=None):
        baseUrl = path or self._baseurl
        if not baseUrl.startswith(self._baseurl):
            return None

        try:
            loader = UrlMetaLoader(baseUrl)
            return ModuleSpec(fullname,
                              loader,
                              is_package=loader.is_package(fullname))
        except Exception as e:
            return None


def install_meta(address):
    finder = UrlMetaFinder(address)
    sys.meta_path.append(finder)


def main():
    install_meta('http://localhost:12800/')
    import temp


if __name__ == "__main__":
    main()
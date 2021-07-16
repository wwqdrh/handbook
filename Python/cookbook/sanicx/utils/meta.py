import copy
import inspect
import typing as T
from importlib import import_module

__all__ = ("func_accepts_kwargs", "import_string", "Proxy")


def func_accepts_kwargs(func):
    return any(
        p for p in inspect.signature(func).parameters.values()
        if p.kind == p.VAR_KEYWORD
    )


def import_string(dotted_path):
    """
    Import a dotted module path and return the attribute/class designated by the
    last name in the path. Raise ImportError if the import failed.
    """
    try:
        module_path, class_name = dotted_path.rsplit('.', 1)
    except ValueError as err:
        raise ImportError("%s doesn't look like a module path" % dotted_path) from err

    module = import_module(module_path)

    try:
        return getattr(module, class_name)
    except AttributeError as err:
        raise ImportError('Module "%s" does not define a "%s" attribute/class' % (
            module_path, class_name)
                          ) from err


class SingletonMeta(type):
    __instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls.__instances:
            cls.__instances[cls] = super().__call__(*args, **kwargs)
        return cls.__instances[cls]

    @property
    def instance(cls):
        return cls.__instances[cls]


class Proxy(T.Protocol):
    @property
    def proxy(self):
        raise NotImplementedError

    __getattr__ = lambda self, item: getattr(self.proxy, item)
    __call__ = lambda self, *args, **kwargs: self.proxy(*args, **kwargs)
    __setattr__ = lambda self, n, v: setattr(self.proxy, n, v)
    __delattr__ = lambda self, n: delattr(self.proxy, n)
    __str__ = lambda self: str(self.proxy)
    __lt__ = lambda self, o: self.proxy < o
    __le__ = lambda self, o: self.proxy <= o
    __eq__ = lambda self, o: self.proxy == o
    __ne__ = lambda self, o: self.proxy != o
    __gt__ = lambda self, o: self.proxy > o
    __ge__ = lambda self, o: self.proxy >= o
    __hash__ = lambda self: hash(self.proxy)
    __len__ = lambda self: len(self.proxy)
    __getitem__ = lambda self, i: self.proxy[i]
    __iter__ = lambda self: iter(self.proxy)
    __contains__ = lambda self, i: i in self.proxy
    __add__ = lambda self, o: self.proxy + o
    __sub__ = lambda self, o: self.proxy - o
    __mul__ = lambda self, o: self.proxy * o
    __floordiv__ = lambda self, o: self.proxy // o
    __mod__ = lambda self, o: self.proxy % o
    __divmod__ = lambda self, o: self.proxy.__divmod__(o)
    __pow__ = lambda self, o: self.proxy ** o
    __lshift__ = lambda self, o: self.proxy << o
    __rshift__ = lambda self, o: self.proxy >> o
    __and__ = lambda self, o: self.proxy & o
    __xor__ = lambda self, o: self.proxy ^ o
    __or__ = lambda self, o: self.proxy | o
    __div__ = lambda self, o: self.proxy.__div__(o)
    __truediv__ = lambda self, o: self.proxy.__truediv__(o)
    __neg__ = lambda self: -self.proxy
    __pos__ = lambda self: +self.proxy
    __abs__ = lambda self: abs(self.proxy)
    __invert__ = lambda self: ~self.proxy
    __complex__ = lambda self: complex(self.proxy)
    __int__ = lambda self: int(self.proxy)
    __long__ = lambda self: long(self.proxy)  # noqa
    __float__ = lambda self: float(self.proxy)
    __oct__ = lambda self: oct(self.proxy)
    __hex__ = lambda self: hex(self.proxy)
    __index__ = lambda self: self.proxy.__index__()
    __coerce__ = lambda self, o: self.proxy.__coerce__(self, o)
    __enter__ = lambda self: self.proxy.__enter__()
    __exit__ = lambda self, *a, **kw: self.proxy.__exit__(*a, **kw)
    __radd__ = lambda self, o: o + self.proxy
    __rsub__ = lambda self, o: o - self.proxy
    __rmul__ = lambda self, o: o * self.proxy
    __rdiv__ = lambda self, o: o / self.proxy
    __rtruediv__ = __rdiv__
    __rfloordiv__ = lambda self, o: o // self.proxy
    __rmod__ = lambda self, o: o % self.proxy
    __rdivmod__ = lambda self, o: self.proxy.__rdivmod__(o)
    __copy__ = lambda self: copy.copy(self.proxy)

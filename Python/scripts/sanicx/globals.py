import contextlib
import typing as T

from sanicx.utils import meta

if T.TYPE_CHECKING:
    from sanicx.app import SanicX
    from sanicx.dependant import ContainerManager

__all__ = ("Proxy", "set_context", "current_app", "container", "bean", "config")


class Proxy(meta.Proxy):
    _context: T.Any = None

    @property
    def proxy(self):
        return object.__getattribute__(self, "_context")

    def set_context(self, context):
        object.__setattr__(self, "_context", context)


Proxy = T.cast(T.Any, Proxy)

current_app: "SanicX" = Proxy()
container: "ContainerManager" = Proxy()
bean: T.Callable = Proxy()
config: T.Callable = Proxy()


def set_context(i, val):
    with contextlib.suppress(Exception):
        context: Proxy = globals()[i]
        context.set_context(val)

from collections import namedtuple, UserDict
from typing import Callable

from sanicx.router.exception_ import RequestValidationError
from sanicx.router.modelutils import Parameter
from sanicx.utils import meta

__all__ = ("RoutesALL",)

Route = namedtuple(
    "Route", ["handler", "methods", "pattern", "parameters", "name", "uri"]
)


class RoutesALL(UserDict):
    IGNORE = ("websocket_handler", "static", "swagger")

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def __setitem__(self, endpoint: str, route: Route):
        if route.name.startswith("websocket_handler"):
            # patch ws_handler 与sanic-doc不兼容问题
            class _wrap(meta.Proxy):
                fn = route.handler

                @property
                def proxy(self):
                    return self.fn

            super().__setitem__(
                endpoint,
                Route(
                    handler=_wrap(),
                    methods=route.methods,
                    pattern=route.pattern,
                    parameters=route.parameters,
                    name=route.name,
                    uri=route.uri,
                ),
            )
        elif any(i in route.name for i in self.IGNORE):
            super().__setitem__(endpoint, route)
        else:
            handler = route.handler
            super().__setitem__(
                endpoint,
                Route(
                    handler=EndPointFunc(handler),
                    methods=route.methods,
                    pattern=route.pattern,
                    parameters=route.parameters,
                    name=route.name,
                    uri=route.uri,
                ),
            )


class EndPointFunc(meta.Proxy):
    """
    代理类
    包装view_func, 这个类会包含原始的参数的类型，然后在调用的时候先检查类型是否正确
    hash eq 用来判断hash值是否相同，类型是否相同，这样索引之类的就是同一个对象
    代理模式
    """

    def __init__(self, fn: Callable):
        object.__setattr__(self, "_fn", fn)
        object.__setattr__(self, "parameter", Parameter.factory(fn))

    @property
    def proxy(self):
        return self._fn

    def __call__(self, request, **kwargs):
        kwargs.update(request=request)
        try:
            json_data = request.json
        except Exception:
            json_data = {}
        values, errors = self.parameter.validate(
            inner_params=kwargs,
            query_params=request.args,
            header_params=request.headers,
            cookie_params=request.cookies,
            body_params=json_data,
            file_params=request.files,
            form_params=request.form,
        )
        if errors:
            raise RequestValidationError(errors, body=request.body)

        kwargs.update(values)
        return super().__call__(**kwargs)  # 函数调用

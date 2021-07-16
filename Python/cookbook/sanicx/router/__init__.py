"""
对于请求参数的校验，以及视图函数相关参数的自动注入
"""
from sanic import response
from sanicx.router.exception_ import RequestValidationError
from sanicx.router.params import (
    Body,
    Cookie,
    File,
    FlaskParam,
    Form,
    Header,
    Param,
    ParamTypes,
    Query,
)
from sanicx.router.router import RoutesALL


def _di_router(app):
    setattr(
        app.router,
        "routes_all",
        RoutesALL(**app.router.routes_all),
    )
    setattr(
        app.router,
        "routes_static",
        RoutesALL(**app.router.routes_static),
    )

    @app.exception(RequestValidationError)
    async def invalid(request, exc):
        return response.json(exc.pretty_errors, status=503)


# [alias]
DIRouter = _di_router

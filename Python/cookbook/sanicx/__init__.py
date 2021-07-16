"""
sanic web 应用脚手架
1、提供参数注入
2、提供依赖注入机制
3、orm、migrate  peewee - TODO
4、更方便的middleware机制
5、提供更方便的帮助函数
6、swagger的配置 使用官方提供的sanic-openapi 需要把ws的去处掉，以及与自定义的参数注入协调 TODO
"""
__version__ = "0.1.0"

from sanic import response

from sanicx import dependant, app, globals
from sanicx.app import *
from sanicx.dependant import *
from sanicx.globals import *

__all__ = (
    response,
    *globals.__all__,
    *app.__all__,
    *dependant.__all__,
)


def get_version() -> str:
    return __version__

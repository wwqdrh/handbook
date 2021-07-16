__all__ = ("SanicX",)

import inspect
import os
import typing as T

from asyncref import executor
from sanic import Sanic
from sanic_openapi import swagger_blueprint
from sanic_session import Session, InMemorySessionInterface
from sanicx import dependant
from sanicx.backend import mysql
from sanicx.core import command
from sanicx.globals import bean, set_context
from sanicx.middleware import IMiddleware
from sanicx.router import DIRouter


class SanicX(Sanic):
    """
    Attributes:
        command: 命令行工具
        extensions: 挂载插件的位置
    Methods:
        builder: 构造函数
        from_cli: 从命令行中运行
        Middleware: 用于注册中间件
    """

    def __new__(self, *args, **kwargs):
        raise NotImplementedError

    def __init__(
            self,
            /,
            name: str,
            router=None,
            error_handler=None,
            load_env=True,
            request_class=None,
            strict_slashes=False,
            log_config=None,
            configure_logging=True,
            session: bool = True,
            swagger_doc: bool = True,
            di_router: bool = True,
            migrate_enable: bool = True,
            container: T.Type[dependant.Container] = None,
            threading_num: int = 4,
    ):
        super().__init__(
            name,
            router,
            error_handler,
            load_env,
            request_class,
            strict_slashes,
            log_config,
            configure_logging,
        )
        executor.max_workers = threading_num  # 设置线程池中最大的线程数
        self.command = command.CommandUtil
        self.container = dependant.ContainerManager.builder(self, container)
        self.extensions = {}

        if swagger_doc:
            self.blueprint(swagger_blueprint)  # swagger开启
        if session:
            Session(self, interface=InMemorySessionInterface())
        if di_router:  # 启动参数注入
            DIRouter(self)
        if migrate_enable:
            mysql.Migrate(self)  # 开启migrate

    @classmethod
    def builder(
            cls,
            /,
            name: str,
            router=None,
            error_handler=None,
            load_env=True,
            request_class=None,
            strict_slashes=False,
            log_config=None,
            configure_logging=True,
            session: bool = True,
            swagger_doc: bool = True,
            di_router: bool = True,
            migrate_enable: bool = True,
            config: T.Union[T.Type, str] = None,
            container: T.Type[dependant.Container] = None,
            threading_num: int = 4,
    ):
        if not (config is None or inspect.isclass(config) or os.path.exists(config)):
            raise Exception(f"给定的路径[{config}]不存在config配置")
        ins = super().__new__(cls)
        ins.__init__(
            name=name,
            router=router,
            error_handler=error_handler,
            load_env=load_env,
            request_class=request_class,
            strict_slashes=strict_slashes,
            log_config=log_config,
            configure_logging=configure_logging,
            session=session,
            swagger_doc=swagger_doc,
            di_router=di_router,
            migrate_enable=migrate_enable,
            container=container,
            threading_num=threading_num,
        )
        if inspect.isclass(config):
            ins.config.from_object(config)
        elif os.path.exists(str(config)):
            ins.config.from_pyfile(config)
        # 全局变量的配置
        set_context("current_app", ins)
        set_context("container", ins.container)
        set_context("bean", lambda i: ins.container.get_bean(i))
        set_context("config", lambda i, default=None: ins.config.get(i, default))

        return ins

    def from_cli(self):
        """
        运行app，检查参数, 执行相应的操作
        subcommand
        """
        self.command().execute()

    def middleware(
            self,
            *uses: T.Union[str, T.Callable, IMiddleware],
            name: str = None,
            attach_to="request",
    ):
        """
        注册中间件，普通函数，包括注册到容器中的中间件函数，或者使用了middleware模板的中间件
        1、middleware模板的中间件：检查如果存在request，那么就注册request相关
        如果存在response相关，就注册response
        """

        def _middleware(handle: T.Callable):
            handle_name = name or handle.__name__
            for use in (bean(use) if isinstance(use, str) else use for use in uses):
                if isinstance(use, IMiddleware):
                    use.register(self, (handle_name,), attach_to=attach_to)
                elif callable(use):
                    self.register_named_middleware(use, (handle_name,), attach_to)
            return handle

        return _middleware

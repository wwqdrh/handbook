"""
1、实现了路由函数功能，给定路由、请求方法定位到视图函数中
2、内置query、form、json请求参数的获取
3、内置以cookie为基础的用户认证模块功能

class MyHandler(HandlerWithAuth):
    who = Router("/who", "get", "who_fn")

    @HandlerWithAuth.auth
    def who_fn(self, username: str):
        return self.send(200, body=bytes(username, encoding="UTF8"))

print("httpserver 8080")
server.HTTPServer(("127.0.0.1", 8080), MyHandler).serve_forever()
"""
from typing import Any, NamedTuple, Callable, Optional, Tuple, Type, ClassVar
from http import server
from urllib import parse
import json
import io
import functools
import string
import random


class Router(NamedTuple):
    """
    url: 路由
    method: 请求方法
    func: 视图函数，以字符串的方式表现，对应handler实例的方法
    """

    url: str
    method: str
    func: str


class Cookie(NamedTuple):
    """
    默认为/全局
    会话 cookie: 将在客户端关闭时被删除。他们没有指定Expires或Max-Age指令。请注意，Web 浏览器通常会启用会话恢复。
    >>> cookie = Cookie(1, 2)
    >>> cookie
    Cookie(name=1, value=2, Path='/' Expires=None, MaxAge=None, HttpOnly=None)
    >>> cookie.to_headers()
    '1=2; Path=/'
    """

    name: str
    value: str
    Path: str = "/"
    Expires: Optional[str] = None
    MaxAge: Optional[int] = None
    HttpOnly: Optional[bool] = None

    def to_headers(self) -> str:
        res = io.StringIO()
        res.write(f"{self.name}={self.value}; Path={self.Path}")
        if self.Expires is not None:
            res.write(f"Expires={self.Expires}; ")
        if self.Expires is not None:
            res.write(f"Max-Age={self.MaxAge}; ")
        if self.HttpOnly is not None:
            res.write(f"HttpOnly")
        res.seek(0)
        return res.read()


class HandlerMeta(server.BaseHTTPRequestHandler):
    """
    加入了基于cookie的认证机制的handler

    /login
    @auth的装饰器，需要判断当前的cookie是否表示某个用户

    >>> class Handler(HandlerMeta):
    ...     home = Router("/", "GET", "homea")
    >>> Handler._router_name
    ('home',)
    """

    _router_name: ClassVar[Tuple[str, ...]]

    do_HEAD = lambda self: self._check_router()
    do_GET = lambda self: self._check_router()
    do_POST = lambda self: self._check_router()
    do_PUT = lambda self: self._check_router()
    do_PATCH = lambda self: self._check_router()
    do_DELETE = lambda self: self._check_router()

    def __init_subclass__(subcls) -> None:
        subcls._router_name = tuple(
            name
            for cls_item in subcls.mro()
            for name, val in cls_item.__dict__.items()
            if isinstance(val, Router)
        )
        return super().__init_subclass__()

    def _check_router(self) -> Any:
        url_parser = parse.urlparse(self.path)
        method = self.command

        found = False
        for router in map(lambda i: getattr(self, i), self._router_name):  # type:Router
            if (
                router.url == url_parser.path
                and router.method.lower() == method.lower()
            ):
                found = True
                try:
                    return getattr(self, router.func)()
                except Exception as e:
                    body = bytes(json.dumps({"message": f"发生异常: {e}"}), encoding="UTF8")
        if not found:
            body = bytes(json.dumps({"message": "没有找到这个路由函数"}), encoding="UTF8")

        return self.send(
            501,
            headers={
                "Content-Type": "application/json",
                "Content-Length": str(len(body)),
            },
            body=body,
        )

    def send(
        self,
        code: int,
        *,
        headers: Optional[dict[str, str]] = None,
        cookies: Optional[tuple[Cookie]] = None,
        body: Optional[bytes] = None,
    ) -> None:
        self.send_response(code)
        if headers is not None:
            for key, value in headers.items():
                self.send_header(key, value)
        if cookies is not None:
            for cookie in cookies:
                self.send_header("Set-Cookie", cookie.to_headers())
        self.end_headers()
        if body is not None:
            self.wfile.write(body)

    def query(self) -> dict[str, list[str]]:
        url_parser = parse.urlparse(self.path)
        args = {}
        for key, val in parse.parse_qs(url_parser.query).items():
            args[key] = list(map(parse.unquote, val))
        return args

    def form(self) -> Optional[dict[str, str]]:
        # TODO: 暂时只写了plain form表单，没有加上文件表单
        if self.headers["Content-Type"] != "application/x-www-form-urlencoded":
            return None

        res: dict[str, Any] = {}

        datas = self.rfile.read(int(self.headers["Content-Length"])).decode("UTF8")
        items = datas.split("&")
        for item in items:
            key, _, val = item.partition("=")
            res[key] = val
        return res

    def json(self) -> Optional[dict[str, Any]]:
        if self.headers["content-type"] != "application/json":
            return None
        datas = self.rfile.read(int(self.headers["content-length"])).decode("UTF8")
        return json.loads(datas)

    def cookie(self) -> dict[str, str]:
        res = {}
        for item in map(lambda i: i.strip(), self.headers["Cookie"].split(";")):
            key, _, val = item.partition("=")
            res[key] = val
        return res


class HandlerWithAuth(HandlerMeta):
    _user_cache: dict[str, str] = {}

    login = Router("/login", "POST", "login_fn")

    @staticmethod
    def auth(fn: Callable):
        """
        视图函数装饰器，表示该视图需要登录了才能访问
        """

        @functools.wraps(fn)
        def _action(ins: HandlerMeta):
            # 判断是否登录了, 读取cookie
            userid = ins.cookie().get("userid", None)
            if userid is None:
                raise Exception("尚未认证")
            return fn(ins, ins._user_cache[userid])  # type: ignore

        return _action

    def login_fn(self):
        """
        登录视图函数
        """
        form = self.form()
        username, password = form["username"], form["password"]
        # 认证机制，这里默认认证成功, 设置cookie
        while True:
            key = "".join(random.sample(string.ascii_letters + string.digits, 16))
            if key not in self._user_cache:
                self._user_cache[key] = username
                break
        self.send(200, cookies=(Cookie(name="userid", value=key),))


if __name__ == "__main__":
    import sys
    import doctest

    if sys.argv[1] == "test":
        doctest.testmod()

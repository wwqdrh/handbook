"""
re match对象group缺省为0，即为整个字符串
在标准库的基础上加上对目录访问的密码限制
需要有注册路由的功能
1、登录页面
2、文件目录访问页面

第一步：
post接收表单的数据，然后
"""
import os
from typing import Type, Dict, Callable, Mapping, ClassVar, Union
from http.server import SimpleHTTPRequestHandler, HTTPServer
from contextlib import contextmanager
import re
import cgi
import argparse
from functools import partial


PAGETEMPLATE = \
"""
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    {head}
    <title>HTTP SERVER</title>
</head>
<body> 
    {body}
<script>
    {script}
</script>
</body>
</html>
"""
INDEXPAGE = ("<h1>index page</h1>")
LOGINPAGE = PAGETEMPLATE.format(head="",
                                body=("""
<div class="container">

    <form class="form-signin" method="POST">
        <h3 class="form-signin-heading">欢迎访问 Flask Http Server</h3>
        <label for="inputName" class="sr-only">用户名</label>
        <input type="text" id="inputName" class="form-control" placeholder="用户名" name="user" required autofocus>
        <label for="inputPassword" class="sr-only">密码</label>
        <input type="password" id="inputPassword" class="form-control" placeholder="密码" name="pwd" required>

        <button class="btn btn-lg btn-primary btn-block" id="submit" type="submit">Login</button>

        <p class="text-danger" style="display:none" id="error"></p>
    </form>

</div>
"""),
                                script=("""
<script>
    window.onload = function(){

    };

    function login_submit(){
        let userName = document.getElementById("inputName").text();
        console.log(userName);
    };
</script>
"""))


class BoostServer(HTTPServer):

    class __SessionManage:
        __id: int = 0
        __current_id: Union[bool, int] = False
        __data: Dict[int, Dict] = {}

        def __setitem__(self, key, val):
            self.__current_id = self.__current_id or self.get_id()
            self.__data.setdefault(self.__current_id, {}).update({key: val})

        @contextmanager
        def cur_context(self, current_id):
            current_id = current_id or self.get_id()
            self.__current_id = current_id
            yield self
            self.__current_id = False

        def get_id(self):
            if not self.__current_id:
                self.__id += 1
                return self.__id
            else:
                return self.__current_id

        def get(self, key: str, id=None):
            if not (id := id or self.__current_id): return None
            session = self.__data.get(id, None)
            if session:
                return session.get(key, None)

    class __Handler(SimpleHTTPRequestHandler):
        outter: ClassVar["BoostServer"]
        __urlMapping: ClassVar[Dict[str, Dict]] = {}
        errorText: ClassVar[str] = "<h1>no this page</h1>"
        methodErrorText: ClassVar[str] = "<h1>不支持这种方法</h1>"

        @property
        def cur_request(self):
            res = {}
            if self.command == "POST":
                form = cgi.FieldStorage(
                    fp=self.rfile,
                    headers=self.headers,
                    environ={
                        'REQUEST_METHOD': 'POST',
                        'CONTENT_TYPE': self.headers['Content-Type'],
                    })
                res["form"] = form
            res["method"] = self.command
            return res

        @classmethod
        def register(cls, url: str, call: Callable, method: list):
            cls.__urlMapping[url] = {"func": call, "method": method}

        def __before_request(self):
            """
            如果没有session就必须跳到登录页面, 如果是注册了的路由就不需要跳转
            登录的get、post或者是注册了的路由，否则就使用默认的方法
            # 没有注册：有没有session
            # 注册了：使用注册方法
            """
            sessionIDMatch = re.search(r"_sessionID=(.*)",
                                        self.headers.get("Cookie", ""))
            sessionID = bool(sessionIDMatch) and sessionIDMatch.group(1)
            user = BoostServer.session.get("user", id=sessionID)
            if (funcInfo := self.__urlMapping.get(self.path, None)) is None \
                or (self.command not in funcInfo["method"]):
                if not user:
                    # 跳转到login
                    self.outter.redirect("/login")(self)
                    return True
                
                return False  # 使用默认的do_GET

            if callable(func := funcInfo["func"]):
                with BoostServer.session.cur_context(sessionID):
                    func(self.cur_request)(self)
                return True
            
            return False

        def do_GET(self):
            if self.__before_request() is False:
                return super().do_GET()

        def do_POST(self):
            if self.__before_request() is False:
                return super().do_POST()

    allow_reuse_address = True
    session: ClassVar["__SessionManage"] = __SessionManage()

    def __init__(self,
                 host: str,
                 port: int,
                 directory: str,
                 handler: Type[__Handler] = __Handler):
        handler.outter = self
        super().__init__((host, port), partial(handler, directory=directory))
        self._host = host
        self._port = port
        self._directory = directory

    def register(self, url: str, method=["GET"]):

        def _register(call: Callable):
            self.__Handler.register(url, call, method)
            return call    # 保留原来的签名因为这里只需要把函数注册进去就行

        return _register

    def html_wrap(self, data: str):

        def cb(handler: SimpleHTTPRequestHandler):
            handler.send_response(200, message="ok")
            handler.send_header("Content-Type", "text/html")
            handler.send_header("Set-Cookie",
                                f"_sessionID={self.session.get_id()}")
            handler.end_headers()
            handler.wfile.write(
                PAGETEMPLATE.format(head="", body=data,
                                    script="").encode(encoding="UTF8"))

        return cb

    def redirect(self, url: str):

        def cb(handler: SimpleHTTPRequestHandler):
            handler.send_response(302, message="redirect login")
            handler.send_header("Location", f"http://{self._host}:{self._port}{url}")
            handler.send_header("Set-Cookie",
                                f"_sessionID={self.session.get_id()}")
            handler.end_headers()
            handler.wfile.write(b"redirection ...")

        return cb


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--directory',
                        '-d',
                        default=os.getcwd(),
                        help='Specify alternative directory '
                        '[default:current directory]')
    parser.add_argument('port',
                        action='store',
                        default=8000,
                        type=int,
                        nargs='?',
                        help='Specify alternate port [default: 8000]')
    args = parser.parse_args()

    server = BoostServer("localhost", args.port, args.directory)

    @server.register("/login", method=["GET", "POST"])
    def login(request):
        if request["method"] == "POST":
            if request["form"].getvalue("user") == "admin" \
                and request["form"].getvalue("pwd") == "123456":
                server.session["user"] = "admin"
                return server.redirect("/")
            else:
                return server.redirect("/login")

        return server.html_wrap(LOGINPAGE)

    @server.register("/home")
    def index(request):
        user = server.session.get("user")
        if not user:
            return server.redirect("/login")
        return server.html_wrap(INDEXPAGE)

    server.serve_forever()


if __name__ == "__main__":
    main()
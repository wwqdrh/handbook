"""
@version: 3.9
@title: http 服务端程序

依赖于http.server的实现
"""
from typing import Type, Tuple
from http.server import BaseHTTPRequestHandler, HTTPServer
import json


RESPONSE: str = (
    "HTTP/1.1 200 OK\r\n" "Server: My server\r\n\r\n" "<h1>Python HTTP Test</h1>"
)


class _HTTPHandler(BaseHTTPRequestHandler):
    TODO = ["You get the response!!"]

    def do_GET(self):
        if self.path != "/":
            self.send_error(404, "Page not Found!")
            return

        resp = json.dumps(self.TODO)
        self.send_response(200, message="OK")
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        self.wfile.write(resp.encode())


class HTTPServerClient(HTTPServer):
    def __init__(
        self,
        address: Tuple[str, int],
        handler: Type[BaseHTTPRequestHandler] = _HTTPHandler,
    ):
        super().__init__(address, handler)


if __name__ == "__main__":
    if len(argv := __import__("sys").argv) > 1 and argv[1] == "test":
        import doctest

        doctest.testmod()

    if len(argv) > 1 and argv[1] == "start":
        http = HTTPServerClient(("localhost", 8080))
        print("Server is running, user <Ctrl+C> to srop.")
        try:
            http.serve_forever()
        finally:
            http.server_close()
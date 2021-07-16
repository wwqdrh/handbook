"""
包括普通的服务器以及ASGI服务器
需要注意的是http服务器当发回了数据之后，这个套接字必须关闭，这样浏览器才不会一直转圈，因为http1.1是不能长连接
1、使用高层库http
HTTPServer 是 socketserver.TCPServer 的一个子类。它会创建和侦听 HTTP 套接字，并将请求调度给处理程序
BaseHTTPRequestHandler此类用于处理到达服务器的HTTP请求。就其本身而言，它不能响应任何实际的HTTP请求。
必须将其子类化以处理每种请求方法（例如GET或POST）。 
BaseHTTPRequestHandler提供了许多类和实例变量，以及供子类使用的方法。
2、使用底层socket库实现阻塞服务器
3、使用selectors多路复用实现非阻塞服务器
3、使用asyncio实现异步服务器
"""
from http.server import BaseHTTPRequestHandler, HTTPServer
from typing import Tuple, Type
import json
import socket
import selectors
from concurrent.futures import ProcessPoolExecutor
import os
import signal
import asyncio

RESPONSE: str = ("HTTP/1.1 200 OK\r\n"
                 "Server: My server\r\n\r\n"
                 "<h1>Python HTTP Test</h1>")


class HTTPServerClient(HTTPServer):

    class __HTTPHandler(BaseHTTPRequestHandler):
        TODO = ['You get the response!!']

        def do_GET(self):
            if self.path != '/':
                self.send_error(404, "Page not Found!")
                return

            resp = json.dumps(self.TODO)
            self.send_response(200, message='OK')
            self.send_header('Content-Type', 'application/json')
            self.end_headers()
            self.wfile.write(resp.encode())

    def __init__(self,
                 address: Tuple[str, int],
                 handler: Type[BaseHTTPRequestHandler] = __HTTPHandler):
        super().__init__(address, handler)


class SocketServerClient:
    """
    socket 基础echo服务器
    """

    def __init__(self, address: Tuple):
        self.address = address

    @classmethod
    def __base_handler(cls, client_socket):
        requestData = client_socket.recv(1024)
        if requestData:
            # 构造响应数据
            response_start_line = "HTTP/1.1 200 OK\r\n"
            response_headers = "Server: My server\r\n"
            response_body = "<h1>Python HTTP Test</h1>"
            response = response_start_line + response_headers + "\r\n" + response_body

            # 向客户端返回响应数据
            client_socket.send(bytes(response, "utf-8"))

        # 关闭客户端连接
        # 关闭一半或全部的连接。如果 how 为 SHUT_RD，则后续不再允许接收。
        # 如果 how 为 SHUT_WR，则后续不再允许发送。
        # 如果 how 为 SHUT_RDWR，则后续的发送和接收都不允许。
        client_socket.shutdown(socket.SHUT_RDWR)
        client_socket.close()

    def serve_forever(self):
        serverSocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        serverSocket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        serverSocket.bind(self.address)
        serverSocket.listen(128)
        while True:
            clientSocket, clientAddress = serverSocket.accept()
            self.__base_handler(clientSocket)


class SelectorServer:

    def __init__(self, address: Tuple):
        self.address = address
        self.selector = selectors.DefaultSelector()

    def __accept(self, sock, mask):

        def read(conn, mask_):
            data = conn.recv(1024)    # Should be ready
            if data:
                conn.send(bytes(RESPONSE, "utf-8"))    # Hope it won't block
            else:
                print('closing', conn)
            self.selector.unregister(conn)
            conn.close()

        conn, addr = sock.accept()    # Should be ready
        conn.setblocking(False)
        self.selector.register(conn, selectors.EVENT_READ, read)

    def serve_forever(self):
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        sock.bind(self.address)
        sock.listen(100)
        sock.setblocking(False)
        self.selector.register(sock, selectors.EVENT_READ,
                               self.__accept)    # read表示可读或者有新的连接加入的情况

        while True:
            events = self.selector.select()    # 选取准备好的事件
            for key, mask in events:
                key.data(key.fileobj, mask)

        self.selector.unregister(sock)
        self.selector.close()


class AsyncioServer:

    def __init__(self, address: Tuple):
        self.address = address

    async def handler(self, reader: asyncio.StreamReader,
                writer: asyncio.StreamWriter):
        data = await reader.read(1024)
        writer.write(bytes(RESPONSE, encoding="UTF8"))
        await writer.drain()

        writer.close()  # 关闭流和基础套接字
        await writer.wait_closed()  # 等待关闭完成

    def serve_forever(self):

        async def _run():
            server = await asyncio.start_server(self.handler, *self.address)
            addr = server.sockets[0].getsockname()
            print(f"Serving on {addr}")
            async with server:
                await server.serve_forever()                                 

        asyncio.run(_run())


def main():
    # http = MyHTTPServerClient(("localhost", 8080))
    # print('Server is running, user <Ctrl+C> to srop.')
    # http.serve_forever()

    # http = SocketServerClient(("localhost", 8080))
    # print('Server is running, user <Ctrl+C> to srop.')
    # http.serve_forever()

    # http = SelectorServer(("localhost", 8080))
    # print('Server is running, user <Ctrl+C> to srop.')
    # http.serve_forever()

    http = AsyncioServer(("localhost", 8080))
    print('Server is running, user <Ctrl+C> to srop.')
    http.serve_forever()


if __name__ == "__main__":
    main()
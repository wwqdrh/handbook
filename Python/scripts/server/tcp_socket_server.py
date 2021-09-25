"""
@version: 3.9
@title: socket 服务端程序

# 关闭客户端连接
# 关闭一半或全部的连接。如果 how 为 SHUT_RD，则后续不再允许接收。
# 如果 how 为 SHUT_WR，则后续不再允许发送。
# 如果 how 为 SHUT_RDWR，则后续的发送和接收都不允许。
"""

from typing import Tuple
import socket


class SocketServerClient:
    """
    socket 基础echo服务器
    """

    def __init__(self, address: Tuple):
        self.address = address

    @classmethod
    def _base_handler(cls, client_socket):
        requestData = client_socket.recv(1024)
        if requestData:
            # 构造响应数据
            response_start_line = "HTTP/1.1 200 OK\r\n"
            response_headers = "Server: My server\r\n"
            response_body = "<h1>Python HTTP Test</h1>"
            response = response_start_line + response_headers + "\r\n" + response_body

            # 向客户端返回响应数据
            client_socket.send(bytes(response, "utf-8"))
        else:
            client_socket.shutdown(socket.SHUT_RDWR)
            client_socket.close()

    def serve_forever(self):
        serverSocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        serverSocket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        serverSocket.bind(self.address)
        serverSocket.listen(128)
        while True:
            clientSocket, clientAddress = serverSocket.accept()
            self._base_handler(clientSocket)

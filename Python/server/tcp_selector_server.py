"""
@version: 3.9
@title: selector 服务端程序


selectors 注册事件 在有相应事件的时候进行回调

- EVENT_READ: 可读
- EVENT_WRITE: 可写

BaseSelector: 用来在多个文件对象上等待 I/O 事件就绪。 它支持文件流注册、注销，以及在这些流上等待 I/O 事件的方法。 它是一个抽象基类，因此不能被实例化。 请改用 DefaultSelector
DefaultSelector: 默认的选择器类，使用当前平台上可用的最高效实现。 这应为大多数用户的默认选择。
SelectSelector: 基于 select.select() 的选择器。
PollSelector: 基于 select.poll() 的选择器。
EpollSelector: 基于 select.epoll() 的选择器。
DevpollSelector: 基于 select.devpoll() 的选择器。
KqueueSelector: 基于 select.kqueue() 的选择器。
"""
from typing import Tuple
import selectors
import socket

RESPONSE: str = (
    "HTTP/1.1 200 OK\r\n" "Server: My server\r\n\r\n" "<h1>Python HTTP Test</h1>"
)


class SelectorServer:
    def __init__(self, address: Tuple):
        self.address = address
        self.selector = selectors.DefaultSelector()

    def _accept(self, sock, mask):
        conn, addr = sock.accept()  # Should be ready
        print("accepted", conn, "from", addr)
        conn.setblocking(False)
        self.selector.register(conn, selectors.EVENT_READ, self._read)

    def read(self, conn, mask):
        data = conn.recv(1024)  # Should be ready

        if data:
            print("echoing", repr(data), "to", conn)
            conn.send(data)  # Hope it won't block
            # conn.send(bytes(RESPONSE, "utf-8"))  # Hope it won't block
        else:
            print("closing", conn)
            self.selector.unregister(conn)
            conn.close()

    def serve_forever(self):
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        sock.bind(self.address)
        sock.listen(100)
        sock.setblocking(False)
        self.selector.register(
            sock, selectors.EVENT_READ, self._accept
        )  # read表示可读或者有新的连接加入的情况
        
        try:
            while True:
                events = self.selector.select()  # 选取准备好的事件
                for key, mask in events:
                    callback = key.data
                    callback(key.fileobj, mask)
        finally:
            self.selector.unregister(sock)
            self.selector.close()

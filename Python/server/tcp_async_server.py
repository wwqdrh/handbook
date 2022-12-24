"""
@version: 3.9
@title: tcp 服务端程序

包含同步 多路复用 异步的编写方式

异步的编写方式 除了 asyncio.start_server的方式外还有其他实现protocol方式然后create_server

# Protocol 关于服务协议部分
- connection_made(): 链接建立时被调用。
- connection_lost(): 链接丢失或关闭时被调用。
- pause_writing(): 传输的缓冲区超过高位标记位时被调用。
- resume_writing(): 传输的缓冲区传送到低位标记位时被调用。
- 流协议
    - data_received(): 接收到数据时被调用。
    - eof_received(): 接收到EOF时被调用。
- 缓冲区协议
    - get_buffer(): 调用后会分配新的接收缓冲区。
    - buffer_updated(): 用接收的数据更新缓冲区时被调用。
    - eof_received(): 接收到EOF时被调用。
- 数据报协议
    - datagram_received(): 接收到数据报时被调用。
    - error_received(): 前一个发送或接收操作引发 OSError 时被调用。
- 子进程协议
    - pipe_data_received(): 子进程向 stdout 或 stderr 管道写入数据时被调用。
    - pipe_connection_lost(): 与子进程通信的其中一个管道关闭时被调用。
    - process_exited(): 子进程退出时被调用。
"""

import asyncio


class EchoServer:
    class _protocol(asyncio.Protocol):
        def connection_made(self, transport: asyncio.Transport):  # type: ignore
            peername = transport.get_extra_info("peername")
            print("Connection from {}".format(peername))
            self.transport = transport

        def data_received(self, data: bytes):
            message = data.decode()
            print("Data received: {!r}".format(message))

            print("Send: {!r}".format(message))
            self.transport.write(data)

            print("Close the client socket")
            self.transport.close()

    async def start(self, host: str, port: int, cb: asyncio.Event):
        loop = asyncio.get_running_loop()
        server = await loop.create_server(self._protocol, host, port)
        async with server:
            cb.set()
            try:
                await server.serve_forever()
            except asyncio.CancelledError:
                print("程序终止")


class EchoServer2:
    """
    tcp server 的第二种实现方式
    """

    async def _handle_echo(
        self, reader: asyncio.StreamReader, writer: asyncio.StreamWriter
    ) -> None:
        data = await reader.read(100)
        message = data.decode()
        addr = writer.get_extra_info("peername")

        print(f"Received {message!r} from {addr!r}")

        print(f"Send: {message!r}")
        writer.write(data)
        await writer.drain()

        print("Close the connection")
        writer.close()

    async def start(self, host: str, port: int, event: asyncio.Event) -> None:
        server = await asyncio.start_server(self._handle_echo, host, port)

        # addr = server.sockets[0].getsockname()
        # print(fServing on {addr})
        # with contextlib.suppress(KeyboardInterrupt):
        async with server:
            event.set()
            try:
                await server.serve_forever()
            except asyncio.CancelledError:
                print("程序终止")
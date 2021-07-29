"""
异步的TCP服务器与客户端实现方式
1、p2p连接，服务端与客户端能够相互通信，收到的消息能够打印出来
"""
from __future__ import annotations

import typing as T
import sys
import asyncio
import argparse
import contextlib


HOST = "127.0.0.1"
PORT = 8888


def _get_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-t", dest="type", choices=["server", "server2", "client"], required=True
    )
    return parser.parse_args(sys.argv[1:])


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

    async def start(self, host: str, port: int):
        loop = asyncio.get_running_loop()
        server = await loop.create_server(self._protocol, HOST, PORT)
        async with server:
            await server.serve_forever()


class EchoClient:
    class _protocol(asyncio.Protocol):
        def __init__(self, message: str, on_con_lost: asyncio.Future):
            self.message = message
            self.on_con_lost = on_con_lost

        def connection_made(self, transport: asyncio.Transport):  # type: ignore
            transport.write(self.message.encode())
            print("Data sent: {!r}".format(self.message))

        def data_received(self, data: bytes):
            print("Data received: {!r}".format(data.decode()))

        def connection_lost(self, exc: None | Exception):
            print("The server closed the connection")
            self.on_con_lost.set_result(True)

    async def start(self, host: str, port: int):
        loop = asyncio.get_running_loop()
        on_con_lost = loop.create_future()
        transport, protocol = await loop.create_connection(
            lambda: self._protocol("Hello", on_con_lost), HOST, PORT
        )

        try:
            await on_con_lost
        finally:
            transport.close()


class StreamServer:
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

    async def start(self, host: str, port: int) -> None:
        server = await asyncio.start_server(self._handle_echo, host, port)

        # addr = server.sockets[0].getsockname()
        # print(f'Serving on {addr}')
        # with contextlib.suppress(KeyboardInterrupt):
        async with server:
            await server.serve_forever()


async def main() -> None:
    args = _get_args()
    if args.type == "server":
        await EchoServer().start(HOST, PORT)
    elif args.type == "server2":
        await StreamServer().start(HOST, PORT)
    elif args.type == "client":
        await EchoClient().start(HOST, PORT)


if __name__ == "__main__":
    with contextlib.suppress(KeyboardInterrupt):
        asyncio.run(main())
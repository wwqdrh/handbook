"""
@version: 3.9
@title: tcp 客户端程序

新建protocol 并 create_connection() -> transport protocol
"""
import asyncio
from typing import Optional


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

        def connection_lost(self, exc: Optional[Exception]):
            print("The server closed the connection")
            self.on_con_lost.set_result(True)

    async def start(self, host: str, port: int):
        loop = asyncio.get_running_loop()
        on_con_lost = loop.create_future()
        transport, protocol = await loop.create_connection(
            lambda: self._protocol("Hello", on_con_lost), host, port
        )

        try:
            await on_con_lost
        finally:
            transport.close()

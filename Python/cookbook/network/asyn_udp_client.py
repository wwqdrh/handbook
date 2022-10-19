import asyncio
from asyncio import DatagramProtocol
import collections


class EchoClientProtocol(DatagramProtocol):
    """
    继承自asyncio的udp协议类
    """

    def __init__(self):
        self.message_queue = collections.deque()
        self.has_message = asyncio.Event()
        self.is_closed = False

    # [override]
    def connection_made(self, transport):
        self.transport = transport

    # [override]
    def datagram_received(self, data, addr):
        self.message_queue.append(data.decode())
        self.has_message.set()

    # [override]
    def error_received(self, exc):
        print("Error received:", exc)
        self.close()

    # [override]
    def connection_lost(self, exc):
        print("Connection closed")
        self.close()

    async def send(self, message: str):
        if self.is_closed:
            raise ValueError("protocol is closed")
        self.transport.sendto(message.encode())

    async def recv(self):
        if self.is_closed:
            raise ValueError("protocol is closed")
        if self.message_queue:
            return self.message_queue.popleft()
        await self.has_message.wait()
        message = self.message_queue.popleft()
        self.has_message.clear()
        return message

    def close(self):
        self.transport.close()
        self.is_closed = True


async def main():
    # Get a reference to the event loop as we plan to use
    # low-level APIs.
    loop = asyncio.get_running_loop()
    transport, protocol = await loop.create_datagram_endpoint(
        lambda: EchoClientProtocol(),
        remote_addr=("127.0.0.1", 9999),
    )

    message = "Hello World!"
    for i in range(1, 6):
        await protocol.send(f"{i}: {message}")
        print(await protocol.recv())
    # print(await protocol.recv())
    # print(await protocol.recv())

    protocol.close()
    # transport.close()


if __name__ == "__main__":
    asyncio.run(main())

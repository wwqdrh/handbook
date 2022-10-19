import asyncio
import signal
from asyncio import DatagramProtocol


class EchoServerProtocol(DatagramProtocol):
    """
    继承自asyncio的udp协议类
    """

    # [override]
    def connection_made(self, transport):
        # 当服务器连接建立完成时
        self.transport = transport

    # [override]
    def datagram_received(self, data, addr):
        # 接受到客户端发来的数据
        message = data.decode()
        print("Received %r from %s" % (message, addr))
        print("Send %r to %s" % (message, addr))
        self.transport.sendto(data, addr)

    # [override]
    def connection_lost(self, exc):
        # 当服务器连接关闭时
        print(f"server close {exc}")


async def main():
    print("Starting UDP server")

    # Get a reference to the event loop as we plan to use
    # low-level APIs.
    loop = asyncio.get_running_loop()
    on_server_shut = loop.create_future()
    signal.signal(signal.SIGINT, lambda *_: on_server_shut.set_result(True))
    # 等待ctrl+c程序退出信号
    # One protocol instance will be created to serve all
    # client requests.
    transport, protocol = await loop.create_datagram_endpoint(
        lambda: EchoServerProtocol(), local_addr=("127.0.0.1", 9999)
    )

    try:
        await on_server_shut
    finally:
        transport.close()


if __name__ == "__main__":
    asyncio.run(main())

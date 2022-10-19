import asyncio
import uvloop


async def echo_server(loop, address):
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.bind(address)
    sock.listen(5)  # 可以同时监听5个
    sock.setblocking(False)  # 设置为非阻塞
    with sock:
        while True:
            client, addr = await loop.sock_accept(sock)
            print("Connection from", addr)
            loop.create_task(echo_client(loop, client))


async def echo_client(loop, client):
    """
    当有一个连接建立成功之后的操作
    """
    try:
        client.setsockopt(socket.IPPROTO_TCP, socket.TCP_NODELAY, 1)
    except (OSError, NameError):
        pass

    with client:
        while True:
            data = await loop.sock_recv(client, 1000000)
            if not data:
                break
            await loop.sock_sendall(client, data)

    print("Connection closed")


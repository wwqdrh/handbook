"""
测试tcpserver async block selector
udpserver ...
"""
import asyncio
import contextlib

import pytest

from scripts.server.tcp_async_server import EchoServer, EchoServer2
from scripts.server.tcp_async_client import EchoClient


HOST = "127.0.0.1"
PORT = 8888


@pytest.mark.asyncio
async def test_server_tcp_async():
    """
    测试tcp异步server部分
    """
    init_event = asyncio.Event()
    server_task = asyncio.create_task(
        EchoServer().start(HOST, PORT, init_event), name="tcp_server"
    )
    await init_event.wait()
    await EchoClient().start(HOST, PORT)

    with contextlib.suppress(asyncio.CancelledError):
        server_task.cancel()


@pytest.mark.asyncio
async def test_server_tcp_async2():
    """
    测试tcp异步server部分
    """
    init_event = asyncio.Event()
    server_task = asyncio.create_task(
        EchoServer2().start(HOST, PORT, init_event), name="tcp_server"
    )
    await init_event.wait()
    await EchoClient().start(HOST, PORT)

    with contextlib.suppress(asyncio.CancelledError):
        server_task.cancel()

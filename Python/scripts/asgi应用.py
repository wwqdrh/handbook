"""
基于ASGI服务器uvicorn的应用实例，uvicorn支持HTTP/1.1, websockets协议
"""

from typing import Awaitable
import io

import uvicorn


class Application:
    async def __call__(self, scope: dict, receive: Awaitable, send: Awaitable):
        assert scope["type"] == "http"

        body = await self.body(receive)  # 根据content-type就能知道是什么类型body，比如表单或者json
        await send(
            {
                "type": "http.response.start",
                "status": 200,
                "headers": [
                    [b"content-type", b"text/plain"],
                ],
            }
        )
        await send(
            {
                "type": "http.response.body",
                "body": b"Hello, world!",
            }
        )

    async def body(self, receive) -> bytes:
        body = io.BytesIO()

        more_body = True
        while more_body:
            message = await receive()
            body.write(message.get("body", b""))
            more_body = message.get("more_body", False)

        body.seek(0)
        return body.read()


if __name__ == "__main__":
    uvicorn.run(Application(), port=8080)
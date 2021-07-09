from fastapi import FastAPI
from fastapi.responses import JSONResponse
from tortoise.contrib.fastapi import register_tortoise
from fastapi.middleware.cors import CORSMiddleware
from fastapi.middleware.gzip import GZipMiddleware

origins = [
    "http://localhost:8080",
]


class Application(FastAPI):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.plugins = {}

    @classmethod
    def run(cls, *args, **kwargs):
        from pyadmin.context import current_app

        ins = cls(*args, **kwargs)
        ins.load_router()
        ins.load_plugins()
        ins.load_middleware()
        ins.load_exception()

        token = current_app.set(ins)
        __import__("uvicorn").run(ins, host="localhost", port=8080, log_level="info")
        current_app.reset(token)

    def load_router(self):
        from pyadmin.apps.api.v1 import api_v1

        self.include_router(api_v1)

    def load_plugins(self):
        from pyadmin.plugins.cache import register_aioredis
        from pyadmin.config import DB_ORM

        register_aioredis(self)
        register_tortoise(self, config=DB_ORM)

    def load_middleware(self):
        self.add_middleware(
            CORSMiddleware,
            allow_origins=origins,
            allow_credentials=True,
            allow_methods=["*"],
            allow_headers=["*"],
        )
        self.add_middleware(GZipMiddleware, minimum_size=1000)

    def load_exception(self):
        self.exception_handler(Exception)

        async def base_exce(request, exc):
            return JSONResponse(content={"msg": "未知错误"}, status_code=501)
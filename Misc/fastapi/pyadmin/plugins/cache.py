"""
async def main():
    redis = aioredis.from_url("redis://localhost")
    await redis.set("key", "string-value")
    bin_value = await redis.get("key")
    assert bin_value == b"string-value"
    redis = aioredis.from_url("redis://localhost", decode_responses=True)
    str_value = await redis.get("key")
    assert str_value == "string-value"

    await redis.close()

"""
from typing import Optional, TYPE_CHECKING

import aioredis

from pyadmin.config import REDIS
from pyadmin.context import Application


def register_aioredis(
    app: Application,
) -> None:
    """
    挂载redispool插件，为app上
    """
    # redis_pool: Optional[ConnectionsPool] = None

    @app.on_event("startup")
    async def init_cache() -> None:  # pylint: disable=W0612
        # nonlocal redis_pool
        redis_pool = await aioredis.create_pool(
            ("localhost", 6379), password=123456, minsize=5, maxsize=10
        )

        app.plugins["redis_pool"] = redis_pool
        # setattr(app, "redis_pool", redis_pool)
        # app.redis_pool = redis_pool

    @app.on_event("shutdown")
    async def close_cache() -> None:  # pylint: disable=W0612
        redis_pool = app.plugins.get("redis_pool", None)
        if redis_pool is not None:
            redis_pool.close()
            await redis_pool.wait_closed()

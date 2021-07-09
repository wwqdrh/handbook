import asyncio

import aioredis


async def main():
    redis_pool = await aioredis.create_pool(
        ("localhost", 6379), password=123456, minsize=5, maxsize=10
    )
    await redis_pool.execute("set", "my-key", "value")
    value = await redis_pool.execute("get", "my-key")
    # async with redis_pool.get() as client:
    #     await client.execute("set", "my-key", "value")

    print(value)
    redis_pool.close()
    await redis_pool.wait_closed()


asyncio.run(main())
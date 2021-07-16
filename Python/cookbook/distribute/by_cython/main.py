import hello
import asyncio
import time


print(hello.primes(10))

def test_async():
    async def main():
        async def t():
            i = 1
            for item in range(10000):
                i += item

        await asyncio.gather(*[t() for i in range(10_000)])

    start_time = time.time()
    asyncio.run(main())
    print(f"python 100000协程 1...10000 执行时间为 {time.time() - start_time}")

    start_time = time.time()
    asyncio.run(hello.main())
    print(f"cython 100000协程 1...10000 执行时间为 {time.time() - start_time}")



test_async()
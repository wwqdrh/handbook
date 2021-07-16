import asyncio

"""
asyncio.cancel
asyncio.future future对象，描述一个异步对象
asyncio.gather 执行两个异步任务并等待他们完成
"""


async def cancel_test():
    async def cancel_me():
        print("cancel_me(): before sleep")

        try:
            await asyncio.sleep(3600)
        except asyncio.CancelledError:
            print("cancel_me(): cancel sleep")
            raise
        finally:
            print("cancel_me(): after sleep")

    task = asyncio.create_task(cancel_me())
    # Wait for 1 second
    await asyncio.sleep(1)

    try:
        task.cancel()  # 比如如果在其他地方调用了这个task.cancel
        await task
    except asyncio.CancelledError:
        print("main(): cancel_me is cancelled now")


async def future_test():
    loop = asyncio.get_running_loop()
    fut = loop.create_future()

    async def set_after(delay, value):
        # Sleep for *delay* seconds.
        await asyncio.sleep(delay)

        # Set *value* as a result of *fut* Future.
        fut.set_result(value)

    loop.create_task(set_after(1, "... world"))
    print("hello ...")
    print(await fut)


async def main():
    await cancel_test()
    await future_test()


asyncio.run(main())
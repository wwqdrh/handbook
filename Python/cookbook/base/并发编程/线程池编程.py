from concurrent import futures
import asyncio

"""
线程池对象的基类：futures._base.Executor
上下文协议, 当退出线程池的时候会对其进行关闭回收资源
def __enter__(self):
    return self

def __exit__(self, exc_type, exc_val, exc_tb):
    self.shutdown(wait=True)
    return False  # 返回false表示异常还是需要外部处理

支持加入同步事件给线程池调度, 
会返回future对象表示这个任务的执行结果
把同步任务给线程池运行避免同步阻塞
"""

# 同步事件
def synchro_submit():
    def run(*args):
        print("向线程池中提交同步事件 {}".format(args))
        return "同步返回成功"

    with futures.ThreadPoolExecutor(max_workers=2) as pool:
        future = pool.submit(run, 1, 2, 3)
        for future_done in futures.as_completed((future,)):
            # 一个生成器，当有某个任务完成，直接打印它的结果
            print(future.result())


async def async_submit():
    # 把同步任务给线程池运行避免同步阻塞
    loop = asyncio.get_running_loop()

    def run(*args):
        print("把同步任务给线程池运行避免同步阻塞 {}".format(args))
        raise Exception("async test exception")
        return "异步返回成功"

    with futures.ThreadPoolExecutor(max_workers=2) as executor:
        future = loop.run_in_executor(executor, run, 1, 2, 3)
        try:
            print(await future)
        except Exception:
            print("出现异常")


def main():
    synchro_submit()
    asyncio.run(async_submit())


if __name__ == "__main__":
    main()

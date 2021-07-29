"""
一个实现了异步事件处理的线程池
"""

from typing import Any, Optional, TYPE_CHECKING, Awaitable
import asyncio
import contextvars
from concurrent import futures
from concurrent.futures import thread
import threading
import weakref

__all__ = ("AsyncExecutor",)

if TYPE_CHECKING:
    from fastapi import FastAPI

_LOOP: contextvars.ContextVar[asyncio.AbstractEventLoop] = contextvars.ContextVar(
    "_LOOP"
)

_EXECUTOR: "_DaemonThreadPoolExecutor"  # 单例模式


def _get_executor() -> "_DaemonThreadPoolExecutor":
    try:
        executor = _EXECUTOR
    except NameError:
        print("executor未作为插件初始化")
        raise RuntimeError("executor未作为插件初始化")
    else:
        return executor


class _DaemonThreadPoolExecutor(thread.ThreadPoolExecutor):
    """
    自定义的线程池类，重写了_adjust_thread_count方法，使得新建的线程为守护线程

    loop描述符，选出一个事件循环
    """

    def _adjust_thread_count(self):
        """
        @override
        """
        # if idle threads are available, don't spin new threads
        if self._idle_semaphore.acquire(timeout=0):
            return

        # When the executor gets lost, the weakref callback will wake up
        # the worker threads.
        def weakref_cb(_, q=self._work_queue):
            q.put(None)

        num_threads = len(self._threads)
        if num_threads < self._max_workers:
            thread_name = "%s_%d" % (self._thread_name_prefix or self, num_threads)
            t = threading.Thread(
                name=thread_name,
                target=thread._worker,
                daemon=True,
                args=(
                    weakref.ref(self, weakref_cb),
                    self._work_queue,
                    self._initializer,
                    self._initargs,
                ),
            )  # 将线程设置为守护线程
            t.start()
            self._threads.add(t)
            thread._threads_queues[t] = self._work_queue

    @property
    def loop(self) -> asyncio.AbstractEventLoop:
        """
        线程池函数
        在异步线程池中选出一个事件循环
        """
        if (loop := _LOOP.get(None)) is None:
            loop = asyncio.new_event_loop()
            asyncio.set_event_loop(loop)
            _LOOP.set(loop)
            self.submit(loop.run_forever)
        return loop

    def execute(self, coro: Awaitable) -> tuple[Any, Optional[Exception]]:
        future = asyncio.run_coroutine_threadsafe(coro, self.loop)
        try:
            result = future.result(5)
        except asyncio.TimeoutError:
            future.cancel()
            return None, Exception("执行超时")
        except Exception as exc:
            return None, exc
        else:
            return result, None


class AsyncExecutor:
    """
    用于fastapi应用的插件，注册线程池的实例化以及回收
    """

    def __init__(self, app: Optional[FastAPI] = None):
        if app is not None:
            self.init_app(app)

    @classmethod
    def async_to_sync(cls, coro: Awaitable) -> Any:
        result, exc = _get_executor().execute(coro)
        if exc is not None:
            raise exc
        return result

    def init_app(self, app: FastAPI):
        @app.on_event("startup")
        async def startup():
            # 创建异步线程池
            global _EXECUTOR
            _EXECUTOR = _DaemonThreadPoolExecutor()
            # 一个用来定义事件循环以及接收任务，一个用来启动事件循环，因为启动之后阻塞无法接收任务所以需要在辅助函数中执行

        @app.on_event("shutdown")
        async def shutdown():
            # 回收异步线程池
            print("回收异步线程")
            _EXECUTOR.shutdown(wait=False)

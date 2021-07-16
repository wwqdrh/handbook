import asyncio
import asyncio
import contextvars
import typing as T
from concurrent import futures

from asyncref import executor

__all__ = ("wrap", "IAdapter")

CoroutineFunc = T.Callable[..., T.Coroutine]
Loop: contextvars.ContextVar[
    T.Optional[asyncio.AbstractEventLoop]
] = contextvars.ContextVar("Loop")
Task: contextvars.ContextVar[asyncio.Queue] = contextvars.ContextVar("Task")


class IAdapter(T.Protocol):
    fn: T.Callable
    args: T.Tuple
    kwargs: T.Dict

    def __call__(self, *args, **kwargs):
        raise NotImplementedError


class SyncAdapter(IAdapter):
    """
    同步代码，在线程池中可以直接执行
    """

    def __init__(self, fn: T.Callable, *args, **kwargs):
        self.fn = fn
        self.args = args
        self.kwargs = kwargs

    def __call__(self, *args, **kwargs):
        return self.fn(*self.args, **self.kwargs)


class AsyncAdapter(IAdapter):
    """
    协程函数转换器，将协程函数包装传给executor执行
    在这个线程中如果不存在事件循环就新建一个
    """

    def __init__(self, fn: T.Union[T.Coroutine, CoroutineFunc], *args, **kwargs):
        self.fn = fn
        self.args = args
        self.kwargs = kwargs

    def __call__(self, *args, **kwargs):
        if (loop := Loop.get(None)) is None:
            loop = asyncio.new_event_loop()
            asyncio.set_event_loop(loop)
            Loop.set(loop)
        try:
            if asyncio.iscoroutine(self.fn):
                return asyncio.run_coroutine_threadsafe(self.fn, loop).result(2)
            else:
                # return asyncio.run_coroutine_threadsafe(self.fn(*self.args, **self.kwargs), loop).result(2)
                return loop.run_until_complete(self.fn(*self.args, **self.kwargs))
        except asyncio.TimeoutError:
            print("执行超时")


def wrap(fn: T.Callable, *args, **kwargs):
    """
    判断fn这个可调用对象是协程函数还是普通函数，
    如果是协程函数就转换为AsyncAdapter
    如果是普通阻塞函数就转换为SyncAdapter
    """
    if asyncio.iscoroutinefunction(fn) or asyncio.iscoroutine(fn):
        # 修改为在主线程执行
        if (loop := Loop.get(None)) is None:
            loop = asyncio.get_event_loop()
            Loop.set(loop)
        if loop.is_closed() or loop.is_running():
            # 主线程的已经在执行了或者说主线程事件循环已经关闭了
            result: futures.Future = executor.executor.submit(
                AsyncAdapter(fn, *args, **kwargs)
            )
            return result.result()
        if asyncio.iscoroutinefunction(fn):
            return loop.run_until_complete(fn(*args, **kwargs))
        else:
            return loop.run_until_complete(fn)
    elif callable(fn):
        loop = asyncio.get_running_loop()
        return loop.run_in_executor(executor.executor, SyncAdapter(fn, *args, **kwargs))

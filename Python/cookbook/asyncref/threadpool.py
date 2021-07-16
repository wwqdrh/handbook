"""
自定义线程池对象，支持通过线程标识符将任务加到指定线程
每个线程中都支持执行异步函数或者同步函数

线程池支持调节线程数量

thread_pool = ThreadPool()

async def a():
    return "hello"

thread_pool.submit(a)
"""
import asyncio
import contextvars
import hashlib
import inspect
import inspect
import logging
import threading
import typing as T
from queue import Queue

__all__ = ("wrap", "Result", "current_thread", "execute")

ThreadShutdown = object()  # 线程终止标识符
_MAX_WORKER = 16


class Task(T.NamedTuple):
    fn: T.Callable
    args: T.Tuple
    kwargs: T.Dict

    def __hash__(self):
        return hash(id(self))


class Result(T.NamedTuple):
    thread_name: str
    result: T.Any


def current_thread() -> str:
    # 当前运行线程的标识
    return threading.current_thread().getName()


class WorkThread(threading.Thread):
    """
    自定义线程类
    定义如何启动，如何终止，如何传递任务
    """

    def __init__(self, thread_pool: "ThreadPool", daemon: bool = True):
        super(WorkThread, self).__init__(daemon=daemon)

        self._thread_pool = thread_pool
        self._task_queue: Queue = Queue()  # 线程池 --函数--> 线程
        self._task_wait: T.Dict[Task, threading.Event] = {}
        self._task_res: T.Dict[Task, T.Any] = {}  # 线程  --结果--> 线程池

    @property
    def is_busy(self) -> bool:
        return len(self._task_wait) != 0

    def add_task(self, task: Task):
        # --主线程中
        # 有些任务可能先执行完成，所以直接使用队列会导致顺序不对
        self._task_queue.put(task)
        event = threading.Event()
        self._task_wait[task] = event

        event.wait()
        self._task_wait.pop(task)
        return self._task_res.pop(task)

    def run(self):
        # --辅助线程中
        # TODO 线程的执行函数逻辑, 如何将函数的计算结果传递出去
        async def _run():
            call: T.Union[ThreadShutdown, Task]

            while True:
                try:
                    if (call := self._task_queue.get()) is ThreadShutdown:
                        break
                    fn, args, kwargs = call.fn, call.args, call.kwargs
                    if asyncio.iscoroutinefunction(fn):
                        res = await fn(*args, **kwargs)
                    elif callable(fn):
                        res = fn(*args, **kwargs)
                        if inspect.isawaitable(res):
                            res = await res

                    self._task_res[call] = res
                    self._task_wait[call].set()
                    self._task_queue.task_done()
                except Exception as e:
                    print(e)

        asyncio.run(_run())


class ThreadPool:
    """
    线程池对象，需要定义如何寻找线程，如何动态调整线程数量，如何将任务传递过去

    1、寻找线程, Dict[str, thread], 创建线程之后根据标识符来寻找指定线程
    2、动态调整线程数量，如果不够就直接添加，如果多了就将已经没有任务的线程回收掉
    3、传递任务，如果指定了线程标识符，那么就传入其中，
        如果没有指定：就遍历一个未开始执行的线程加入进去，如果没有且当前线程数小于最大值那么就可以新建一个，如果已经到饱和了就将任务加入到第一个
        这里的submit，提交之后可以立即得到结果，
    """

    def __init__(self, min_thread_num: int = 1, max_thread_num: int = 16):
        self._min_thread_num = max(1, min_thread_num)
        self._max_thread_num = min(16, max_thread_num)
        self._cur_thread_num = self._min_thread_num

        self._map_thread = {}
        for _ in range(self._cur_thread_num):
            thread = WorkThread(self)
            thread.start()
            self._map_thread[thread.name] = thread

    def _adjust_thread(self) -> T.Optional[WorkThread]:
        if self._cur_thread_num < self._max_thread_num:
            thread = WorkThread(thread_pool=self, daemon=True)
            self._cur_thread_num += 1
            self._map_thread[thread.name] = thread
            return thread
        else:
            return None

    def submit(
            self,
            task: T.Callable,
            args: T.Tuple = None,
            kwargs: T.Dict = None,
            thread_name: str = "",
    ) -> Result:
        """
        传递任务
        """
        thread: T.Optional[WorkThread]
        args = args or tuple()
        kwargs = kwargs or dict()
        if thread_name != "" and (thread := self._map_thread.get(thread_name, None)):
            return Result(
                thread_name=thread.name,
                result=thread.add_task(Task(task, args, kwargs)),
            )
        else:
            # 判断空闲的线程加入，如果没有且当前小于最大就新建一个，否则就加入到随机一个
            for thread in self._map_thread.values():
                if thread.is_busy:
                    continue
                return Result(
                    thread_name=thread.name,
                    result=thread.add_task(Task(task, args, kwargs)),
                )
            thread = self._adjust_thread() or next(iter(self._map_thread.values()))
            return Result(
                thread_name=thread.name,
                result=thread.add_task(Task(task, args, kwargs)),
            )


executor: ThreadPool = ThreadPool(max_thread_num=_MAX_WORKER)


def wrap(fn: T.Callable, args: T.Tuple = None, kwargs: T.Dict = None, thread_name: str = "") -> Result:
    """
    判断fn这个可调用对象是协程函数还是普通函数，
    如果是协程函数就转换为AsyncAdapter
    如果是普通阻塞函数就转换为SyncAdapter
    """
    return executor.submit(fn, args=args, kwargs=kwargs, thread_name=thread_name)


def execute(thread_name: str = ""):
    def _execute(fn):
        return wrap(fn, thread_name=thread_name)

    return _execute

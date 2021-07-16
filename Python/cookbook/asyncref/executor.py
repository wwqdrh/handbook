"""
线程池对象，重写线程池，需要能够动态调整线程池中线程的数量，需要能够指定在某个线程中执行代码
"""
import typing as T
from concurrent import futures

__all__ = ("executor",)


_MAX_WORKER = 1


class AdjustExecutor(futures.ThreadPoolExecutor):
    """
    可调节最大worker的线程池
    """

    _max_workers: int

    def __init__(
        self, max_workers=None, thread_name_prefix="", initializer=None, initargs=()
    ):
        super(AdjustExecutor, self).__init__(
            max_workers, thread_name_prefix, initializer, initargs
        )

    @property
    def max_workers(self) -> int:
        return self._max_workers

    @max_workers.setter
    def max_workers(self, worker: int):
        # 设置了_max_workers之后需要调整
        self._max_workers = worker
        self._adjust_thread_count()  # type: ignore

    def __del__(self):
        self.shutdown(wait=True)


executor: "AdjustExecutor" = AdjustExecutor(max_workers=_MAX_WORKER)



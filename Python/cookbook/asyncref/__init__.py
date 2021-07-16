"""
1、事件循环外执行协程函数，协程函数中将阻塞操作转为协程对象
>>
维护一个全局线程池对象，将协程函数或者普通函数传入给转换器
转换器会采用相应的策略来转换对象，使其能够直接提交给线程池中进行执行
"""
from .threadpool import *

__all__ = threadpool.__all__
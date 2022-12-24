"""
N个产品经理提出任务需求，M个程序员根据规律来选择完成某个需求
计算每个任务所需要的时间

选取规则：
    1、先选择优先级高的
    2、优先级相同选择所需时间小的idea
    3、时间相同选择PM序号最小的
"""
from dataclasses import dataclass, field
import heapq
from typing import *
from utils.queue import PriorityQueue


@dataclass
class Task:
    # 优先级高 所需时间小 PM序号小
    PMId: int
    startTime: int
    priority: int
    needTime: int
    finishTime: int = field(init=False, default=None)

    def __str__(self):
        return "[%s] | 开始时间: %s | 需要时间: %s | 结束时间: %s" \
            % (self.PMId, self.startTime, self.needTime, self.finishTime)


class Scheduler:
    """
    维护一个优先级队列
    """

    def __init__(self):
        self.task = []    # 存储添加的任务的顺序，方便之后打印
        self.createQ = PriorityQueue(
            lambda i: i.startTime)    # 优先级队列，生产者，按照创建时间进行排序，最小堆
        self.taskQ = PriorityQueue(
            lambda i:
            (i.needTime, i.PMId))    # 优先级队列, 消费者， 构造一个最大堆，每次程序员寻找最大的任务

    def submit(self, task: Task):
        # 添加任务
        self.task.append(task)
        self.createQ.offer(task)
        # heapq.heappush(self._createQueue, -task)

    def run(self, coder: int):
        createQ, taskQ = self.createQ, self.taskQ
        program = [0] * coder    # 表示每一个程序员还有多少时间有空
        timer = 0    # 计时器
        while createQ or taskQ:
            timer += 1
            while createQ and createQ.peek(
            ).startTime == timer:    # 将当前时间段的数据加入到任务队列中
                taskQ.offer(createQ.poll())
            for idx in range(coder):
                if program[idx] > 0:
                    program[idx] -= 1
                else:
                    if not taskQ:
                        continue
                    task = taskQ.poll()
                    task.finishTime = timer + task.needTime
                    program[idx] = task.needTime - 1    # 表示用户获取这个任务之后执行了一秒
        for task in self.task:
            print(task)


def main():
    scheduler = Scheduler()
    scheduler.submit(Task(1, 1, 1, 2))
    scheduler.submit(Task(1, 2, 1, 1))
    scheduler.submit(Task(1, 3, 2, 2))
    scheduler.submit(Task(2, 1, 1, 2))
    scheduler.submit(Task(2, 3, 5, 5))
    scheduler.run(2)


if __name__ == "__main__":
    main()
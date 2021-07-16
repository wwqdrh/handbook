from typing import *
import heapq

__all__ = ["PriorityQueue"]

class PriorityQueue:

    def __init__(self, call):
        self.call = call
        self.data = []

    def __bool__(self):
        return bool(self.data)

    def peek(self):
        if not self.data:
            return
        return self.data[0][-1]

    def poll(self):
        if not self.data:
            return
        *_, data = heapq.heappop(self.data)
        return data

    def offer(self, data):
        key = self.call(data)
        heapq.heappush(self.data, (key, id(data), data))

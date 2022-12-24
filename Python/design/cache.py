"""
定义与缓存相关的算法
"""
from typing import Dict, List


class LFUCache:
    """
    LFU Cache 最不经常使用缓存
    思路：与lrucache不同的是，这个需要为每一个访问的键进行计数，另外如果有一些键的访问次数相同，
    那么就需要使用一个linked来将上一次访问时间最久的那一个进行删除
    put(key, value) - 如果键不存在，请设置或插入值。
    当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。
    在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最远 最少使用的键。
    「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。
    请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。它应该支持以下操作：get 和 put。
    get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
    """
    def __init__(self, capacity: int):
        # 会更新每个元素出现的频率，当频率相同的时候删除最久未使用
        self.capacity = capacity
        self._data: Dict = {}  # 存储数据, key: 频率
        self._freq: List[Dict] = [{}]  # 存储每一个频率所包含的key
        
    def get(self, key: int) -> int:
        # 将频率为n的key转到频率为n+1的linked中
        if key not in self._data: return -1
        if (freq := self._data[key]) == len(self._freq):
            self._freq.append({})
        self._data[key] += 1
        val = self._freq[freq-1].pop(key)
        self._freq[freq].update({key: val})
        return val

    def put(self, key: int, value: int) -> None:
        if self.capacity <= 0: return
        if (freq := self._data.get(key, 0)):  # 如果存在这个key
            self.get(key)
            self._freq[freq][key] = value
            return
        
        if len(self._data) == self.capacity:
            minFreq = next(i for i in self._freq if len(i) != 0)
            leftKey = next(iter(minFreq.keys()))
            minFreq.pop(leftKey)
            self._data.pop(leftKey)
        self._data[key] = 1
        self._freq[0].update({key: value})

class LRUCache:
    def __init__(self, capacity: int):
        self._capactity = capacity
        self._data: Dict = {}

    def get(self, key: int) -> int:
        if key not in self._data: return -1
        val = self._data.pop(key)
        self._data[key] = val
        return val

    def put(self, key: int, value: int) -> None:
        if self._capactity <= 0: return
        if self.get(key) != -1:  # 存在
            self._data[key] = value
            return
        # 容量满了
        if len(self._data) == self._capactity:
            oldKey = next(iter(self._data.keys()))
            self._data.pop(oldKey)
        self._data[key] = value
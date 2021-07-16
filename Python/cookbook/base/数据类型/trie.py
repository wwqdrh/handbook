"""
字典树的定义以及常用方法
"""
from typing import List, Iterable, Callable
from dataclasses import dataclass, field, InitVar


@dataclass
class TrieTree:
    val: int = field(default=-1)
    isLeaf: bool = field(init=False, repr=False, default=False)
    child: List = field(init=False, repr=False)
    capacity: InitVar[int] = 0

    def __post_init__(self, capacity):
        self.child = [None] * capacity    # 定义初始化容量

    @staticmethod
    def build_trie(datas: List[Iterable], capacity: int, get_idx: Callable):
        """
        根据给定容器以及给定大小构造字典树
        """
        root = TrieTree(capacity=capacity)
        for data in datas:
            node = root
            for item in data:
                idx = get_idx(item)
                if node.child[idx] is None:
                    node.child[idx] = TrieTree(val=item, capacity=capacity)
                    node = node.child[idx]
            node.isLeaf = True
        return root


def intDictOrder(max_num: int, order_idx: int, start: int = 1):
    """
    给定整数范围内按照字典序排序第order_idx项
    默认是1到end，暂时不让指定start

    可用字典序树的方法来做，但并不用生成字典序树。
    从第一个分支开始，计算该分支的节点数目num，若order_idx > num,
    跳到下一个分支，同时end −= order_idx，依次类推。
    若order_idx<num,说明r要求的数在该分支中，跳到该分支的第一个子分支。
    """

    cur = 1    # 当前节点
    order_idx -= 1
    while order_idx > 0:
        num = 0    # 当前分支的元素个数
        start, end = cur, cur + 1
        while start <= max_num:    # 保证开始元素要比max_num小，避免越界
            num += min(max_num + 1, end) - start    # 计数当前分支有多少个元素了
            start *= 10
            end *= 10
        if num > order_idx:    # 说明要找的元素就在这个分支中
            cur *= 10
            order_idx -= 1
        else:
            cur += 1
            order_idx -= num
    return cur


if __name__ == "__main__":
    # 测试字符串字典树的构建
    tree = TrieTree.build_trie(["looked", "just", "like", "her", "brother"],
                               capacity=26,
                               get_idx=lambda i: ord(i) - 97)
    print(tree)

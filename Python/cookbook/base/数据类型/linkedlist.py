"""
定义了链表的结构ListNode
以及针对链表的插入排序，归并排序
针对链表的归并排序
时间复杂度O(nlogn) 空间复杂度O(1)
"""
from typing import *
from dataclasses import dataclass


@dataclass
class ListNode:
    val: int
    next_: "ListNode" = None

    def __str__(self) -> str:
        res = []
        cur = self
        while cur:
            res.append(str(cur.val))
            cur = cur.next_
        return "->".join(res)

    @property
    def length(self) -> int:
        res = 0
        cur = self
        while cur:
            cur = cur.next_
            res += 1
        return res

    @classmethod
    def construct(cls, nums: List[int]):
        dummyHead = ListNode(-1)
        cur = dummyHead
        for num in nums:
            cur.next_ = ListNode(num)
            cur = cur.next_
        return dummyHead.next_


def merge_sort(head: ListNode) -> ListNode:
    # 自底向上归并排序
    factor, length = 1, head.length    # factor: 因子
    res = ListNode(-1, next_=head)
    while factor < length:
        pre, h = res, res.next_    # 每次使用res.next_才能使得在新的节点中
        while h:
            h1, c1 = h, factor    # 第一部分的head，以及factor-c1个元素
            while h and c1:
                h, c1 = h.next_, c1 - 1
            if c1 > 0:
                break
            h2, c2 = h, factor    # 第二部分的head，以及factor-c2个元素
            while h and c2:
                h, c2 = h.next_, c2 - 1
            c1, c2 = factor - c1, factor - c2    # 长度
            while c1 or c2:    # 并操作
                val1 = h1.val if c1 > 0 else float("inf")
                val2 = h2.val if c2 > 0 else float("inf")
                if val1 < val2:
                    pre.next_, h1 = h1, h1.next_
                    c1 -= 1
                else:
                    pre.next_, h2 = h2, h2.next_
                    c2 -= 1
                pre = pre.next_
            pre.next_ = h    # 将排序了的和未排序的接上
        factor *= 2
    return res.next_


def insertionSortList(head: ListNode) -> ListNode:
    """
    插入排序是当前元素与它左边的进行比较，看是否需要插入到合适的位置
    
    在遍历的时候，需要保存前一个元素的引用，以及最左边的引用
    """
    left = ListNode(-1)
    left.next_ = head
    while head and head.next_:
        if head.val <= head.next_.val:
            head = head.next_
            continue
        pre = left
        while pre.next_.val < head.next_.val:
            pre = pre.next_
        curr = head.next_
        pre.next_, curr.next_, head.next_ = curr, pre.next_, curr.next_
    return left.next_


if __name__ == "__main__":
    res = merge_sort(ListNode.construct([1, -5, 3, 4, 0]))
    # res = insertionSortList(ListNode.construct([1, -5, 3, 4, 0]))
    print(str(res))
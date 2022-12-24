"""
关于图的一系列方法
"""

from typing import List, Dict, Set


def neighbor_table_graph(nums: List[List[int]], start: int, target: int):
    """
    传递参数构造邻接表表示方法, 然后判断两个元素是否存在路径
    [[0, 1], [0, 2], [1, 2], [1, 2]]
    :nums 给定一个列表包含图中的两个元素的连接情况，判断开始位置和终点位置是否是通路
    """

    def _dfs(cur: int) -> bool:
        if cur == target:  # 找到目标节点
            return True
        if cur in visited or cur not in graph_:  # 已经访问过了或者不是图中的节点
            return False  
        visited.add(cur)
        for neigh in graph_[cur]:
            if _dfs(neigh):
                return True
        visited.remove(cur)
        return False

    # 构造图
    graph_: Dict[int, List[int]] = {}
    for cur, neigh in nums:
        graph_.setdefault(cur, []).append(neigh)

    # 深搜判断元素是否存在路径
    visited: Set = set()
    return _dfs(start)
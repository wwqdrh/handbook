from typing import Dict
from enum import Enum

TreeType = Enum('TreeType', 'apple_tree cherry_tree peach_tree')


class Tree:
    pool: Dict = dict()

    def __new__(cls, tree_type, *args, **kwargs):
        obj = cls.pool.get(tree_type, None)
        if not obj:
            obj = super().__new__(cls, *args, **kwargs)
            cls.pool[tree_type] = obj
            obj.tree_type = tree_type
        return obj

    def __init__(self, size):
        self.size = size

    def render(self, age, x, y):
        print('render a tree of type {} and age {} at ({},{})'.format(
            self.tree_type, age, x, y))

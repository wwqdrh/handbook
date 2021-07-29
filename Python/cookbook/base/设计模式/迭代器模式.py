
'''
节点类
@class Node 单个节点
@class MyTree 整个节点
'''
class Node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None

class MyTree:
    '''
    @method add_node 从根节点开始遍历直到某个左右节点为空，大的放右边，小的放左边
    @method __iter__ __next__ 实现了迭代器协议，用于迭代节点
    @method getgenerator generator版本，用于简化__iter__ __next__
    '''
    def __init__(self, root):
        self.root = root
    def __iter__(self):
        if self.root is None:
            self.stack = []
        else:
            self.stack = [self.root]
            current = self.root
            while current.left is not None:
                current = current.left
                self.stack.append(current)
        return self
    def __next__(self):
        if len(self.stack) <= 0:
            raise StopIteration
        while len(self.stack) > 0:
            current = self.stack.pop()
            data = current.data
            if current.right is not None:
                current = current.right
                self.stack.append(current)
                while current.left is not None:
                    current = current.left
                    self.stack.append(current)
            return data
        raise StopIteration
    def add_node(self, node):
        current = self.root
        while True:
            if node.data <= current.data:
                if current.left is None:
                    current.left = node
                    return
                else:
                    current = current.left
            else:
                if current.right is None:
                    current.right = node
                    return
                else:
                    current = current.right
    def getgenerator(self):
        if self.root is None:
            self.stack = []
        else:
            self.stack = [self.root]
            current = self.root
            while current.left is not None:
                current = current.left
                self.stack.append(current)
        
        if len(self.stack) <= 0:
            # generator 函数不应该抛出StopIteration
            return

        while len(self.stack) > 0:
            current = self.stack.pop()
            data = current.data
            if current.right is not None:
                current = current.right
                self.stack.append(current)
                while current.left is not None:
                    current = current.left
                    self.stack.append(current)
            yield data
        

if __name__ == "__main__":
    tree = MyTree(Node(16))
    tree.add_node(Node(8))
    tree.add_node(Node(1))
    tree.add_node(Node(17))
    tree.add_node(Node(13))
    tree.add_node(Node(14))
    tree.add_node(Node(9))
    tree.add_node(Node(10))
    tree.add_node(Node(11))

    for i in tree.getgenerator():
        print(i)
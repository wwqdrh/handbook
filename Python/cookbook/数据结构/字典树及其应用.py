from typing import *
"""
用来统计某个文本出现频率最大的词
"""


class TrieNode:

    def __init__(self, var=None, parent=None, num=0):
        self.num = num
        self.isEnd = False
        self.son = {}
        self.var = var
        self.parent = parent


class Trie:

    def __init__(self):
        self.root = TrieNode()

    def insert(self, str):
        if len(str) <= 0:
            return
        node = self.root
        for c in str:
            if c not in node.son.keys():
                node.son[c] = TrieNode(c, node, 1)
            else:
                node.son[c].num += 1
            node = node.son[c]
        node.isEnd = True

    def has(self, str):
        if len(str) == 0:
            return False
        node = self.root
        for c in str:
            if c not in node.son.keys():
                return False
            else:
                node = node.son[c]
        return node.isEnd

    def countPrefix(self, prefix):
        if len(prefix) == 0:
            return -1
        node = self.root
        for c in prefix:
            if c not in node.son.keys():
                return 0
            else:
                node = node.son[c]
        return node.num

    def preOrder(self, node):
        if node != None:
            for child in node.son.keys():
                self.preOrder(node.son[child])

    def mostString(self):
        max = [0]
        r = [TrieNode()]
        self.helper(self.root, max, r)
        x = r[0]
        s = []
        while x != None:
            s.append(x.var)
            x = x.parent
        s.reverse()
        return s

    def helper(self, node, max, r):
        if node != None:
            if node.isEnd and node.num >= max[0]:
                r[0] = node
                max[0] = node.num
            for child in node.son.keys():
                self.helper(node.son[child], max, r)


def mostString2():
    dict = {}
    t = Trie()
    fr = open('preprocessing.txt')
    for line in fr.readlines():
        for s in line.strip().split(' '):
            t.insert(s)
    #print(t.preOrder(t.root))
    #print(t)
    #t.preOrder(t.root)
    #print(k)
    print(t.has('chen'))
    print(t.has('chendsfdsfsd'))
    print(t.countPrefix('chen'))
    print(t.mostString())


"""
给定一个字典以及一个待匹配的字符串，判断当前不能匹配的长度
"""


def origin(dictonary: List[str], sentences: str):
    """
    使用动态规划，dp[i]表示 [0,i)最少的未识别数
    """
    dic = set(dictonary)
    length = len(sentences)
    dp = [0] * (length + 1)
    for i in range(1, length + 1):
        dp[i] = dp[i - 1] + 1
        for j in range(i):
            if sentences[j:i] in dic:
                dp[i] = min(dp[i], dp[j])
    return dp[-1]


"""
上述代码套了两层循环，缺点就是对于每一个i，它前面的子字符串都被找了个遍，
这其中包括一些根本不可能在字典中出现的单词。需要找一个方法提前结束。
一种方法是记录字典中每个单词最后一个字符，如果想匹配的字符串的最后一个字母都不在字典里，
就没必要再看这个字符串了；此外，即使字符串最后一个字母在词典里，也不用挨个去找[j,i)子字符串是否匹配，
即不需要让j从0到i遍历一遍，只要看对应长度的子串在不在词典里即可。

使用字典树进行优化
"""
from dataclasses import dataclass, field


class TrieSolution:

    @dataclass
    class _TrieNode:
        childs: List[Optional["TrieSolution._TrieNode"]] = field(
            init=False, default_factory=lambda: [None] * 26)
        isWord: bool = field(init=False, default=False)    # 表示当前节点结尾是否是一个词

    @classmethod
    def trie_cons(cls, dictonary: List[str]) -> _TrieNode:
        root = cls._TrieNode()
        for word in dictonary:
            node = root
            for char in word:
                idx = ord(char) - 97
                if node.childs[idx] is None:
                    node.childs[idx] = cls._TrieNode()
                node = node.childs[idx]
            node.isWord = True    # 代表单词的结尾
        return root

    def respace(self, dictionary: List[str], sentences: str):
        root = self.trie_cons(dictionary)    # 创建字典树
        n = len(sentences)
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            node = root
            dp[i] = i
            for j in range(i):
                idx = ord(sentences[j]) - 97
                if node.childs[idx] is None:
                    dp[i] = min(dp[i], i - j + dp[j - 1])
                    break
                if node.childs[idx].isWord:
                    dp[i] = min(dp[i], dp[j - 1])
                else:
                    dp[i] = min(dp[i], i - j + dp[j - 1])
                node = node.childs[idx]
        return dp[-1]


"""
给定一个数组，寻找其中元素两两异或的结果大于指定元素的数量
暴力解法会超时，构建字典树
"""


@dataclass
class TrieTree:
    count: int = field(default=1)
    child: List = field(init=False, repr=False)

    def __post_init__(self):
        self.child = [None] * 2


def get_xor_cnt(arras: List[int], target: int):

    def _build_trie():
        # 使用字典树存储列表中元素每一位的元素
        root = TrieTree()
        for num in arras:
            node = root
            for j in range(31, -1, -1):
                digit = (num >> j) & 1    # 从第31位高位开始判断是0还是1
                if node.child[digit] is None:
                    node.child[digit] = TrieTree()
                else:
                    node.child[digit].count += 1
                node = node.child[digit]
        return root

    def _query_trie(trie: TrieTree, ori: int, targ: int, idx: int):
        # 搜索字典树中的对象
        if not trie:
            return 0
        current = trie
        for i in range(idx, -1, -1):
            aDigit = (ori >> i) & 1
            bDigit = (targ >> i) & 1
            if aDigit == 1 and bDigit == 1:
                if current.child[
                        0] is None:    # 字典k位为1，a数k位为1，异或结果为0，目标值为1，说明当前元素不满足条件
                    return 0
                current = current.child[0]
            elif aDigit == 0 and bDigit == 1:
                if current.child[1] is None:    # 异或结果为0，目标为1，不满足
                    return 0
                current = current.child[1]
            elif aDigit == 1 and bDigit == 0:
                p = _query_trie(current.child[1], ori, targ, idx - 1)
                q = 0 if current.child[0] is None else current.child[0].count
                return p + q
            elif aDigit == 0 and bDigit == 0:
                p = _query_trie(current.child[0], ori, targ, idx - 1)
                q = 0 if current.child[1] is None else current.child[1].count
                return p + q
        return 0

    trieTree = _build_trie()
    result = 0
    for num in arras:
        result += _query_trie(trieTree, num, target, 31)
    return result // 2


if __name__ == "__main__":
    print(get_xor_cnt([6, 5, 10], 10))
    # TrieSolution().respace(["looked","just","like","her","brother"], "jesslookedjustliketimherbrother")
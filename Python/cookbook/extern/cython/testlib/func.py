# cython: annotation_typing = False
import time


def flib(n: int) -> int:
    """
    计算从1到a的累加和
    >>> flib(3)
    6
    """
    start = time.time()
    res: int = 0
    for i in range(1, n + 1):
        res += i
    print(f"cost: {time.time() - start}")
    return res


if __name__ == "__main__":
    # import sys

    # if sys.argv[1] == "test":
    import doctest

    doctest.testmod()

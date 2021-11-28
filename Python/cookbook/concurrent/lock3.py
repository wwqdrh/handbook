import threading

lst = []


def even():
    """加偶数"""
    with condLock:
        for i in range(2, 101, 2):
            condLock.wait()  # 交出执行权，等待另一个线程通知加偶数
            lst.append(i)
            condLock.notify()
        condLock.notify()


def odd():
    """加奇数"""
    with condLock:
        for i in range(1, 101, 2):
            # if len(lst) % 2 == 0:
            lst.append(i)
            condLock.notify()
            condLock.wait()
        condLock.notify()


if __name__ == "__main__":
    condLock = threading.Condition()

    addEvenTask = threading.Thread(target=even)
    addOddTask = threading.Thread(target=odd)

    addEvenTask.start()
    addOddTask.start()

    addEvenTask.join()
    addOddTask.join()

    print(lst)

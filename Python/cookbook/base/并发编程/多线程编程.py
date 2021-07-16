import threading
import time
from queue import Queue

"""
多线程实现方法
"""

started_evt = threading.Event()  # 用于控制线程状态的同步事件

# 继承Thread
class MyThread1(threading.Thread):
    def run(self):
        # @Override
        super().run()
        print("{}--{}".format("MyThread", threading.current_thread()))


# 直接使用Thread构造
def MyThread2(daemon=None):
    def countdown(n, started_evt: threading.Event):
        # 用于测试的多线程执行入口
        print("countdown starting {}".format(threading.current_thread()))
        started_evt.set()
        while n > 0:
            print("T-minutes", n)
            n -= 1
            time.sleep(0.5)

    return threading.Thread(target=countdown, daemon=daemon, args=(4, started_evt))


"""
线程间通信，使用线程安全的队列来通信
"""


def concat():
    queue = Queue(maxsize=5)

    def producer():
        data = 0
        for i in range(5):
            data += 1
            queue.put(data)

    def consumer():
        for i in range(5):
            data = queue.get()
            print("the consumer get the data {}".format(data))

    prod = threading.Thread(target=producer)
    cons = threading.Thread(target=consumer)
    prod.start()
    cons.start()
    prod.join()
    cons.join()


def main():
    print("Launching countdown")
    t1 = MyThread1()
    t2 = MyThread2(daemon=True)  # 如果是守护线程，那么当主线程退出的时候守护线程也会退出，默认不是守护线程，这样主程序会等待辅助线程执行完成
    t1.start()
    t2.start()
    started_evt.wait()  # wait for the started_evt start and the next go on
    print("countdown is running")
    concat()


if __name__ == "__main__":
    main()
import multiprocessing
from queue import Queue

"""
multiprocessing 模块还引入了在 threading 模块中没有的API。
一个主要的例子就是 Pool 对象，它提供了一种快捷的方法，赋予函数并行化处理一系列输入值的能力，
可以将输入数据分配给不同进程处理（数据并行）。
"""


class MyProcess1(multiprocessing.Process):
    def run(self):
        # @Override
        super().run()
        print("自定义进程执行\n")


def MyProcess2():
    def run(*args):
        print("自定义进程执行 {}\n".format(args))

    return multiprocessing.Process(target=run, args=(1, 2))


def main():
    process1 = MyProcess1()
    process2 = MyProcess2()
    process1.start()
    process2.start()
    process1.join()  # 等待子进程执行完毕，避免成为僵尸进程
    process2.join()


if __name__ == "__main__":
    main()
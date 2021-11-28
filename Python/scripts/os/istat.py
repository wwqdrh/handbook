"""
@version: 3.9
@title: istat 系统资源分析

需要注意的是可能没有权限执行 需要加上sudo

内存
mem = psutil.virtual_memory()

cpu信息

网络相关

进程管理
"""
import datetime

import psutil


def show_memory():
    mem = psutil.virtual_memory()
    # 系统总计内存
    zj = float(mem.total) / 1024 / 1024 / 1024
    # 系统已经使用内存
    ysy = float(mem.used) / 1024 / 1024 / 1024
    # 系统空闲内存
    kx = float(mem.free) / 1024 / 1024 / 1024
    print("系统总计内存:{:.2f}GB".format(zj))
    print("系统已经使用内存{:.2f}GB".format(ysy))
    print("系统空闲内存:{:.2f}GB".format(kx))


def show_cpu():
    print(psutil.cpu_count())  # 获取逻辑的CPU个数
    print(psutil.cpu_count(logical=False))  # 获取物理的CPU个数
    print(psutil.cpu_percent())  # 获取CPU的利用率
    print(psutil.cpu_percent(percpu=True))  # 获取所有逻辑CPU个数的利用率
    print(psutil.cpu_percent(percpu=True, interval=2))  # 获取2秒内所有逻辑CPU个数的利用率
    print(psutil.cpu_times(percpu=True))  # 查看所有CPU的时间花费
    print(psutil.cpu_percent(percpu=True, interval=3))  # 查看3秒内所有CPU的时间花费的比例
    print(psutil.cpu_stats())  # 查看CPU上下文切换，中断，软中断，系统调用次数


def show_disk():
    print(psutil.disk_partitions())  # 查看磁盘名称、挂载点、文件系统的类型等信息

    # disk = [
    #         item for item in psutil.disk_partitions() if item.mountpoint == "/"
    # ]
    print(psutil.disk_usage("/Users"))  # 查看硬盘的大小，已使用磁盘容量，空间利用率等
    print(psutil.disk_io_counters(perdisk=True))  # 查看每个磁盘的读的次数，写的次数，读字节数，写的字节数等，省去了解


def show_network():
    print(psutil.net_io_counters())  # 获取所有网口的所有流量和包
    print(psutil.net_io_counters(pernic=True))  # 获取每张网口的流量和包
    print(psutil.net_connections())  # 获取每个网络连接的信息
    print(psutil.net_if_addrs())  # 获取网卡的配置信息
    print(psutil.net_if_stats())  # 获取网卡是否启动、通信类型，MTU，传输速度等


def show_spec():
    print(psutil.users())  # 获取用户登陆信息

    print(psutil.boot_time())  # 返回启动系统的时间戳
    t = datetime.datetime.fromtimestamp(psutil.boot_time()).strftime(
        "%Y-%m-%d %H:%M:%S"
    )
    print(t)


def show_process(pid: int):
    if not psutil.pid_exists(pid):  # 判断pid是否存在
        print(f"pid: {pid}: 不存在 ")
        return

    init_process = psutil.Process(pid)  # 获取linux 第一个进程
    print(init_process.cmdline())  # 获取启动的程序位置
    print(init_process.name())  # 获取进程名字
    print(
        datetime.datetime.fromtimestamp(init_process.create_time()).strftime(
            "%Y-%m-%d %H:%M:%S"
        )
    )  # 获取进程启动时间
    print(init_process.num_fds())  # 获取文件打开个数
    # print(init_process.threads())  # 获取子进程的个数  AccessDeny
    print(init_process.is_running())  # 判断进程是否运行着
    # init_process.send_signal(signal.SIGKILL)#发送信号给进程
    # init_process.kill() #发送SIGKILL信号结束进程
    # init_process.terminate()  # 发送SIGTERM信号结束进程

    print(psutil.pids())  # 获取正在运行的所在pid


if __name__ == "__main__":
    if len(argv := __import__("sys").argv) > 1 and argv[1] == "test":
        import doctest

        doctest.testmod()

    if len(argv) > 1:
        command = argv[1]
        if command == "memory":
            show_memory()
        elif command == "cpu":
            show_cpu()
        elif command == "disk":
            show_disk()
        elif command == "network":
            show_network()
        elif command == "spec":
            show_spec()
        elif command == "process" and len(argv) > 2:
            show_process(int(argv[2]))

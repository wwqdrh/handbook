import datetime
import jinja2
import socket

import psutil
import os
import yagmail


def render(tpl_path, *args, **kwargs):
    """渲染html模板"""
    path, filename = os.path.split(tpl_path)
    return (
        jinja2.Environment(loader=jinja2.FileSystemLoader(path or "./"))
        .get_template(filename)
        .render(**kwargs)
    )


def byte2human(n):
    """
    字节自动转为可读性强的函数
    :param n:字节数
    :return:
    """
    symbols = ("K", "M", "G", "T", "P", "E", "Z", "Y")
    prefix = {}
    for i, s in enumerate(symbols):
        prefix[s] = 1 << (i + 1) * 10  # 左位移2位乘10
    for s in reversed(symbols):
        if n >= prefix[s]:  # 只在大于1000的时候，才转换为K
            value = float(n) / prefix[s]
            return "%.1f%s" % (value, s)
    return "%sB" % n


def get_cpu_info():
    """获取CPU个数和1秒内CPU的利用率"""
    cpu_count = psutil.cpu_count()
    cpu_percent = psutil.cpu_percent(interval=1)
    return dict(cpu_count=cpu_count, cpu_percent=cpu_percent)


def get_memory_info():
    """获取内存的信息"""
    virtual_mem = psutil.virtual_memory()
    mem_total = byte2human(virtual_mem.total)  # 总内存大小
    mem_percent = virtual_mem.percent  # 内存的使用率

    mem_free = byte2human(virtual_mem.free)  # 未使用内存
    mem_used = byte2human(virtual_mem.total * (virtual_mem.percent / 100))  # 使用内存
    return dict(
        mem_total=mem_total,
        mem_percent=mem_percent,
        mem_free=mem_free,
        mem_used=mem_used,
    )


def get_disk_info():
    """获取硬盘的信息"""
    disk_usage = psutil.disk_usage("C:\\")
    disk_total = byte2human(disk_usage.total)  # 硬盘总大小
    disk_percent = disk_usage.percent  # 硬盘的使用率
    disk_free = byte2human(disk_usage.free)  # 硬盘未使用大小
    disk_used = byte2human(disk_usage.used)  # 硬盘的使用大小
    return dict(
        disk_total=disk_total,
        disk_percent=disk_percent,
        disk_free=disk_free,
        disk_used=disk_used,
    )


def get_boot_info():
    """查看系统的开始时间"""
    boot_time = datetime.datetime.fromtimestamp(psutil.boot_time()).strftime(
        "%Y-%m-%d %H:%M:%S"
    )
    return dict(boot_time=boot_time)


def collect_monitor_data():
    """将所有字典组合成在一起"""
    data = {}
    data.update(get_boot_info())
    data.update(get_cpu_info())
    data.update(get_disk_info())
    data.update(get_memory_info())
    return data


if __name__ == "__main__":
    email_username = "xxx@qq.com"
    email_password = "xxxx"
    smtp_ip = "smtp.ym.163.com"

    RECV_EMAIL = ["xxx@qq.com"]

    hostname = socket.gethostname()
    data = collect_monitor_data()
    data.update(hostname=hostname)

    content = render("index.html", **data)

    with yagmail.SMTP(email_username, email_password, host=smtp_ip, port=465) as yag:
        for to_email in RECV_EMAIL:
            yag.send(to_email, subject="监控信息", contents=content)

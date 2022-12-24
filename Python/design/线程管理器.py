from concurrent.futures import ThreadPoolExecutor, as_completed
import threading
import time

aDone = threading.Event()
bDone = threading.Event()
cDone = threading.Event()


def entryA():
    # 每一个线程的入口函数
    while True:
        print("A---- %s" % threading.current_thread())
        aDone.set()
        cDone.wait()
        cDone.clear()
        time.sleep(2)


def entryB():
    while True:
        aDone.wait()
        print("B---- %s" % threading.current_thread())
        aDone.clear()
        bDone.set()
        time.sleep(2)


def entryC():
    while True:
        bDone.wait()
        print("C---- %s" % threading.current_thread())
        bDone.clear()
        cDone.set()
        time.sleep(2)


with ThreadPoolExecutor(max_workers=3) as executor:
    threadEntry = [entryA, entryB, entryC]
    threadTask = [executor.submit(call) for call in threadEntry]
    for res in as_completed(threadTask):
        print(res.exception())
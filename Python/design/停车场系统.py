"""
这个管理系统可实现车辆入库，按车牌号或者车型查询车辆，修改车辆信息，
车辆出库时实现计费，按车型统计车辆数和显示全部车辆信息的功能
"""
"""
车
    - 车牌号
    - 车型

车库
    - 按车牌或者车型查询车辆
    - 按车型统计车辆类型信息
    - 显示全部车辆信息
    - 车辆入库操作
    - 车辆出库操作
"""
import time
from dataclasses import dataclass, field
import weakref
from typing import *


@dataclass(frozen=True)  # frozen设置为True那么这个对象是可哈希的
class Car:
    stub: str
    cardId: str

    def __str__(self):
        return "[%s] %s" % (self.stub, self.cardId)


@dataclass(unsafe_hash=True)  # 只需要保证需要hash的字段不变就行了
class CarPoolInfo:
    """
    当有一辆车入库之后包装一个信息
    """
    startTime: str  # hash=None 使用compare的值，默认为True
    endTime: str = field(init=False, hash=False, default="$$")  # 由于
    car: Car = field(compare=False, hash=True)

    def __str__(self):
        return "%s - %s: %s" % (self.startTime, self.endTime, self.car)
    
    def __del__(self):
        print("车出库，资源回收")


class PoolSystem:

    def __init__(self):
        self.__allCar: Dict[str, CarPoolInfo] = {}    # id: car 所有进来的车按序存储在这
        self.__queryStub: Dict[str, weakref.WeakSet] = {}    # 按照车类型来存储

    @property
    def time(self) -> str:
        # 返回当前时间
        localTime = time.localtime(time.time())
        return time.strftime(r"%Y-%m-%d %H:%M:%S", localTime)

    def car_in(self, car: Car):
        # 有一辆车入库
        carInfo = CarPoolInfo(startTime=self.time, car=car)
        self.__allCar[car.cardId] = carInfo
        self.__queryStub.setdefault(car.stub, weakref.WeakSet()).add(carInfo)

    def car_out(self, car: Car):
        # 有一辆车出库
        carInfo = self.__allCar.pop(car.cardId)
        carInfo.endTime = self.time
        print("从%s开始，到%s结束" % (carInfo.startTime, carInfo.endTime))

    def query(self, cardId: str):
        # 根据车牌号进行查询
        if car := self.__allCar.get(cardId, None):
            print(str(car))
        else:
            print("%s 不存在" % cardId)

    def show_stub(self, stub: str):
        # 显示当前停车场stub类型的车辆有哪些
        print("----展示当前类型所有的车辆----")
        if cars := self.__queryStub.get(stub, None):
            for car in cars:
                print(str(car))
        else:
            print("-------无----------")
        print("--------------")

    def all_car(self):
        # 显示当前停车场所有的车
        print("----展示所有的车辆----")
        for car in self.__allCar.values():
            print(str(car))
        print("--------------")

if __name__ == "__main__":
    poolSystem = PoolSystem()

    car1 = Car("奔驰", "A10000")

    poolSystem.car_in(car1)
    poolSystem.query("A10000")
    poolSystem.show_stub("奔驰")
    poolSystem.show_stub("宝马")
    poolSystem.all_car()
    poolSystem.car_out(car1)
    print("-------")
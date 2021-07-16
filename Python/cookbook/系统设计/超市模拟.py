'''
题目：模拟一个超市

对象：
    @class SuperMarket 超市
        @method 超市可以购货
        @method 超市可以展示货物
        @method 超市可以雇佣店员
        @method 超市可以展示店员
    @class Goods 商品
        @property 商品名字
        @property 商品价格
        @property 商品生产日期
    @class ShopAssistant 店员
        @method 店员可以表示是否在忙
        @method 店员可以找零
    @class Customer 顾客
        @method 顾客可以进商店
        @method 顾客可以挑选物品到购物车
        @method 顾客可以结账

设计模式：
    1、商品类：可以使用享元模式，避免过多的new操作
'''
from abc import ABC
from enum import Enum
from typing import Dict, Type, Iterable, List


class GoodsAbstract(ABC):
    def __init__(self, spec: Dict):
        self.name = spec.get('name', None)
        self.price = spec.get('price', None)

    def __str__(self):
        return '---{}---{}'.format(self.name, self.price)


class ToothPaste(GoodsAbstract):
    def __init__(self, spec: Dict):
        initopts = {'name': '牙膏', 'price': 10.02}
        spec = {
            key: spec[key] if key in spec else initopts[key]
            for key in initopts
        }
        super().__init__(spec)


class NoGoodsError:
    instance = None

    def __new__(cls):
        if not cls.instance:
            cls.instance = super().__init__(cls)
        return cls.instance

    def __init__(self):
        self.info = '对不起，没有该商品'

    def __bool__(self):
        return False


class Goods:
    pool: Dict = dict()
    parser_alias = {'牙膏': ToothPaste, 'default': NoGoodsError}

    def __new__(cls, goods_type: str):
        obj = cls.pool.get(goods_type, None)
        if not obj:
            obj = cls.parser_alias.get(goods_type,
                                       cls.parser_alias['default'])({})
            cls.pool[goods_type] = obj
        return obj


class StoreContainer:
    def __init__(self):
        spec = ['生活用品', '食品', '玩具']
        self.goods_spec = {key: dict() for key in spec}

    def add(self, goods_spce_type: str, goods_type: str):
        if goods_spce_type not in self.goods_spec:
            print('{},对不起，没有{}该商品分类'.format('*' * 10, goods_spce_type))
            return

        goods = Goods(goods_type)
        if not goods:
            print('{},对不起，没有{}该商品项目'.format('*' * 10, goods_type))
            return
        
        if goods_type not in self.goods_spec[goods_spce_type]:
            self.goods_spec[goods_spce_type][goods_type] = {
                'goods': goods,
                'num': 1
            }
        else:
            self.goods_spec[goods_spce_type][goods_type]['num'] += 1

    def remove(self, goods_spce_type: str, goods_type: str):
        spec = self.goods_spec
        if goods_spce_type not in spec:
            print('{},对不起，没有{}该商品分类'.format('*' * 10, goods_spce_type))
            return

        if goods_type not in spec[goods_spce_type]:
            print('{},对不起，没有{}该商品项目'.format('*' * 10, goods_type))
            return

        spec[goods_spce_type][goods_type]['num'] -= 1
        return spec[goods_spce_type][goods_type]['goods']

    def show(self):
        print('\n------')
        goods_spec = self.goods_spec
        for spec_type in goods_spec:
            print('{}{}'.format('*' * 10, spec_type))
            for goods in goods_spec[spec_type].values():
                print('{}{}   :剩余量{}'.format('*' * 15, goods['goods'], goods['num']))


class SuperMarket:
    def __init__(self):
        self.store_container = StoreContainer()

    def purchase(self, goods_spce_type: str, goods_type: str):
        self.store_container.add(goods_spce_type, goods_type)

    def goodsTaken(self, goods_spce_type: str, goods_type: str):
        return self.store_container.remove(goods_spce_type, goods_type)

    def showGoods(self):
        self.store_container.show()


class Customer:
    def __init__(self):
        self.shopping_cart = []

    def inDoor(self, supermarket: SuperMarket):
        self.market_place = supermarket

    def viewGoods(self):
        self.market_place.showGoods()

    def showCart(self):
        for item in self.shopping_cart:
            print(item)

    def buyGoods(self, goods_spce_type: str, goods_type: str):
        goods = self.market_place.goodsTaken(goods_spce_type, goods_type)
        if not goods:
            print('没有该商品，请求非法')
        else:
            self.shopping_cart.append(goods)


if __name__ == "__main__":
    # 商场有这些东西
    supermarket = SuperMarket()
    supermarket.purchase('生活用品', '牙膏')
    supermarket.purchase('生活用品', '牙膏')

    # 定义用户
    customer1 = Customer()
    customer1.inDoor(supermarket)
    customer1.viewGoods()
    customer1.buyGoods('生活用品', '牙膏')
    customer1.viewGoods()
    customer1.showCart()

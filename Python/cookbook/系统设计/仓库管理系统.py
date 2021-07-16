"""
有一个手机仓储管理系统，使用者有三方：销售、仓库管理员、采购。

需求是：
销售：
    - 销售一旦达成订单，销售人员会通过系统的销售子系统部分通知仓储子系统，
    - 仓储子系统会将可出仓手机数量减少，同时通知采购管理子系统当前销售订单；
购买：
    - 仓储子系统的库存到达阈值以下，会通知销售子系统和采购子系统，并督促采购子系统采购；
    - 采购完成后，采购人员会把采购信息填入采购子系统，
    - 采购子系统会通知销售子系统采购完成，并通知仓库子系统增加库存。

子系统
- 销售子系统：销售 生成销售订单
- 仓库子系统：存储 售出
- 采购子系统：购买
"""
"""
中介者模式：
从需求描述来看，每个子系统都和其它子系统有所交流，
在设计系统时，如果直接在一个子系统中集成对另两个子系统的操作，一是耦合太大，二是不易扩展。为解决这类问题，
我们需要引入一个新的角色-中介者-来将“网状结构”精简为“星形结构”。（为充分说明设计模式，某些系统细节暂时不考虑，例如：仓库满了怎么办该怎么设计。类似业务性的内容暂时不考虑）
"""


class StockMediator:

    def __init__(self):
        self.sell_system = None
        self.stock_system = None
        self.buy_system = None

    def load_system(self, sell_system, stock_system, buy_system):
        self.sell_system = sell_system
        self.stock_system = stock_system
        self.buy_system = buy_system

    def sell(self, number: int):
        # 售出number件商品
        self.sell_system.sell_order(number)    # 打印订单
        self.stock_system.goods_out(number)    # 库存减少

    def buy(self, number: int):
        # 购买number件商品
        self.buy_system.buy_order(number)    # 打印订单
        self.stock_system.goods_in(number)    # 库存增加


class SellSystem:

    def __init__(self, *, mediator):
        self.mediator = mediator

    def sell_order(self, number: int):
        # 生成销售订单
        order = ("-------------\n"
                 "--- order ---\n"
                 "--- %s件  ---\n"
                 "-------------\n") % number
        print(order)


class StockSystem:

    def __init__(self, threadshold: int, *, mediator):
        self.mediator = mediator
        self._stock = 0
        self._threadshold = threadshold    # 当低于这个线的时候就需要进货了

    def goods_in(self, number: int):
        # 进货
        self._stock += number

    def goods_out(self, number: int):
        # 出货
        self._stock -= number
        if self._stock < self._threadshold:
            print("库存低于最低标准，开始进货100件")
            self.mediator.buy(100)


class BuySystem:

    def __init__(self, *, mediator):
        self.mediator = mediator

    def buy_order(self, number: int):
        order = ("-------------\n"
                 "--- buy   ---\n"
                 "--- %s件  ---\n"
                 "-------------\n") % number
        print(order)


if __name__ == "__main__":
    mediator = StockMediator()
    sell_system = SellSystem(mediator=mediator)
    stock_system = StockSystem(500, mediator=mediator)
    buy_system = BuySystem(mediator=mediator)
    mediator.load_system(sell_system, stock_system, buy_system)

    # 测试
    mediator.buy(100)
    mediator.buy(1000)
    mediator.sell(700)
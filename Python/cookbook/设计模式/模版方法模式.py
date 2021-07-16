'''
模版方法会严格按照其结构体所表明的结构来进行处理。
它提供了一种方法模版，可以遵循该模版逐步实现一个特定过程，
然后，只要简单地修改一两处细节，就可以将该模版
运用到不同的场景里
'''
from abc import ABCMeta, abstractmethod


'''
定义模版类
@class ThirdPartyInteractionTemplate
'''

class ThirdPartyInteractionTemplate(metaclass=ABCMeta):
    # __metaclass__ = ABCMeta # 这种方法只能在python2使用，python3这样用没有效果
    def sync_stock_items(self):
        self._sync_stock_items_step_1()
        self._sync_stock_items_step_2()
        self._sync_stock_items_step_3()
        self._sync_stock_items_step_4()
    def send_transaction(self, transaction):
        self._send_transaction(transaction)
    @abstractmethod
    def _sync_stock_items_step_1(self):
        pass
    @abstractmethod
    def _sync_stock_items_step_2(self):
        pass
    @abstractmethod
    def _sync_stock_items_step_3(self):
        pass
    @abstractmethod
    def _sync_stock_items_step_4(self):
        pass
    @abstractmethod
    def _send_transaction(self, transaction):
        pass

'''
定义具体的实现类
@class System1
@class System2
@class System3
'''
class System1(ThirdPartyInteractionTemplate):
    def _sync_stock_items_step_1(self):
        print('step1 1')
    def _sync_stock_items_step_2(self):
        print('step2 1')
    def _sync_stock_items_step_3(self):
        print('step3 1')
    def _sync_stock_items_step_4(self):
        print('step4 1')
    def _send_transaction(self, transaction):
        print("send transaction to system1: {0!r}".format(transaction))
class System2(ThirdPartyInteractionTemplate):
    def _sync_stock_items_step_1(self):
        print('step1 2')
    def _sync_stock_items_step_2(self):
        print('step2 2')
    def _sync_stock_items_step_3(self):
        print('step3 2')
    def _sync_stock_items_step_4(self):
        print('step4 2')
    def _send_transaction(self, transaction):
        print("send transaction to system2: {0!r}".format(transaction))
class System3(ThirdPartyInteractionTemplate):
    def _sync_stock_items_step_1(self):
        print('step1 3')
    def _sync_stock_items_step_2(self):
        print('step2 3')
    def _sync_stock_items_step_3(self):
        print('step3 3')
    def _sync_stock_items_step_4(self):
        print('step4 3')
    def _send_transaction(self, transaction):
        print("send transaction to system3: {0!r}".format(transaction))

if __name__ == "__main__":
    transaction = {
        'id': 1,
        'items': [
            {
                'item_id': 1,
                'amount_purchased': 3,
                'value': 238
            }
        ]
    }
    for c in [System1, System2, System3]:
        print("="*10)
        system = c()
        system.sync_stock_items()
        system.send_transaction(transaction)
'''
策略模式，
允许编写使用一些策略的代码，以便在运行的时候进行选择，
而除了要遵循执行特征之外，不需要知道关于该策略的任何信息
'''

'''
相对应的策略类
@class AdditionStrategy
@class SubtractionStrategy
@class StrategyExcutor
'''

class AdditionStrategy:
    def execute(self, arg1, arg2):
        print(arg1 + arg2)

class SubtractionStrategy:
    def execute(self, arg1, arg2):
        print(arg1 - arg2)

class StrategyExcutor:
    def __init__(self, strategy=None):
        self.strategy = strategy
    def execute(self, *args):
        if self.strategy is None:
            print('Strategy not implemented...')
        else:
            self.strategy.execute(*args)

if __name__ == "__main__":
    no_strategy = StrategyExcutor()
    addtion_strategy = StrategyExcutor(AdditionStrategy())
    substraction_strategy = StrategyExcutor(SubtractionStrategy())

    no_strategy.execute(4, 6)
    addtion_strategy.execute(4, 6)
    substraction_strategy.execute(4, 6)
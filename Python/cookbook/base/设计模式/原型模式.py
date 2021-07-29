from abc import ABCMeta, abstractmethod
from copy import deepcopy


class Prototype(metaclass=ABCMeta):
    @abstractmethod
    def clone(self):
        pass

class Knight(Prototype):
    def __init__(self, level):
        self.unit_type = 'knight'
        filename = '{}_{}.dat'.format(self.unit_type, level)
        self.info = filename
    def __str__(self):
        return '读取 {} 配置'.format(self.info)
    def clone(self):
        return deepcopy(self)

class Archer(Prototype):
    def __init__(self, level):
        self.unit_type = 'archer'
        filename = '{}_{}.dat'.format(self.unit_type, level)

        self.info = filename
    def __str__(self):
        return '读取 {} 配置'.format(self.info)
    def clone(self):
        return deepcopy(self)

class Barracks:
    def __init__(self):
        self.units = {
            'knight': {
                1: Knight(1),
                2: Knight(2)
            },
            'archer': {
                1: Archer(1),
                2: Archer(2)
            }
        }
    def build_unit(self, unit_type, level):
        return self.units[unit_type][level].clone()

if __name__ == "__main__":
    barrcks = Barracks()
    knight1 = barrcks.build_unit('knight', 1)
    knight2 = barrcks.build_unit('knight', 2)
    print('[knight] {}'.format(knight1))
    print('[knight] {}'.format(knight2))
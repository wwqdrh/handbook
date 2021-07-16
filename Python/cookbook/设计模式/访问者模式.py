import abc
import random
import unittest

'''
定义可访问对象
@metaclass Visitable
@class {Visitable} CompositeVisitable
@class {Visitable} Light
@class {Visitable} Thermostat
@class {Visitable} ThemperatureRegulator
@class {Visitable} DoorLock
@class {Visitable} CoffeeMachine
@class {Visitable} Clock
@class {CompositeVisitable} MyHomeSystem !! 主体类
'''
class Visitable:
    def accept(self, visitor):
        visitor.visit(self)
class CompositeVisitable(Visitable):
    def __init__(self, iterable):
        self.iterable = iterable
    def accept(self, visitor):
        for element in self.iterable:
            element.accept(visitor)
        visitor.visit(self)
class Light(Visitable):
    def __init__(self, name):
        self.name = name
        self.status = self.get_status()
    def get_status(self):
        return random.choice(range(-1, 2))
    def is_online(self):
        return self.status != -1
    def boot_up(self):
        self.status = 0
class Thermostat(Visitable):
    def __init__(self, name):
        self.name = name
        self.status = self.get_status()
    def get_status(self):
        temp_range = [x for x in range(-10, 31)]
        temp_range.append(None)
        return random.choice(temp_range)
    def is_online(self):
        return self.status is not None
    def boot_up(self):
        pass
class ThemperatureRegulator(Visitable):
    def __init__(self, name):
        self.name = name
        self.status = self.get_status()
    def get_status(self):
        return random.choice(['heating', 'cooling', 'on', 'off', 'error'])
    def is_online(self):
        return self.status != 'error'
    def boot_up(self):
        pass
class DoorLock(Visitable):
    def __init__(self, name):
        self.name = name
        self.status = self.get_status()
    def get_status(self):
        return random.choice(range(-1, 2))
    def is_online(self):
        return self.status != -1
    def boot_up(self):
        pass
class CoffeeMachine(Visitable):
    def __init__(self, name):
        self.name = name
        self.status = self.get_status()
    def get_status(self):
        return random.choice(range(-1, 5))
    def is_online(self):
        return self.status != -1
    def boot_up(self):
        self.status = 1
class Clock(Visitable):
    def __init__(self, name):
        self.name = name
        self.status = self.get_status()
    def get_status(self):
        return "{:0>2}:{:0>2}".format(random.randrange(24), random.randrange(60))
    def is_online(self):
        return True
    def boot_up(self):
        self.status = "00:00"

class MyHomeSystem(CompositeVisitable):
    pass

'''
定义相关访问者
@abstractclass AbstractVisitor
@class {AbstractVisitor} CompositeVisitor
@class {AbstractVisitor} MyHomeSystemStatusUpdateVisitor
@class {AbstractVisitor} LightStatusUpdateVisitor
@class {AbstractVisitor} ThermostatStatusUpdateVisitor
@class {AbstractVisitor} TemperatureRegulatorStatusUpdateVisitor
@class {AbstractVisitor} DoorLockStatusUpdateVisitor
@class {AbstractVisitor} CoffeeMachineStatusUpdateVisitor
@class {AbstractVisitor} ClockStatusUpdateVisitor
'''
class AbstractVisitor(abc.ABC):
    @abc.abstractmethod
    def visit(self, element):
        raise NotImplementedError('A visitor need to define a visit method')
class CompositeVisitor(AbstractVisitor):
    def __init__(self, person_1_home, person_2_home):
        self.person_1_home = person_1_home
        self.person_2_home = person_2_home
    def visit(self, element):
        try:
            c = eval("{}StatusUpdateVisitor".format(element.__class__.__name__))
        except:
            print("{}StatusUpdateVisitor not found".format(element.__class__.__name__))
        else:
            visitor = c(self.person_1_home, self.person_2_home)
            visitor.visit(element)
class MyHomeSystemStatusUpdateVisitor(AbstractVisitor):
    def __init__(self, person1_home, person2_home):
        self.person1_home = person1_home
        self.person2_home = person2_home
    def visit(self, element):
        pass
class LightStatusUpdateVisitor(AbstractVisitor):
    def __init__(self, person1_home, person2_home):
        self.person1_home = person1_home
        self.person2_home = person2_home
    def visit(self, element):
        if self.person2_home:
            element.status = 1
        else:
            element.status = 0
class ThermostatStatusUpdateVisitor(AbstractVisitor):
    def __init__(self, person1_home, person2_home):
        self.person1_home = person1_home
        self.person2_home = person2_home
    def visit(self, element):
        pass
class TemperatureRegulatorStatusUpdateVisitor(AbstractVisitor):
    def __init__(self, person1_home, person2_home):
        self.person1_home = person1_home
        self.person2_home = person2_home
    def visit(self, element):
        if self.person1_home:
            if self.person2_home:
                element.status = 'on'
            else:
                element.status = 'heating'
        elif self.person2_home:
            element.status = 'cooling'
        else:
            element.status = 'off'
class DoorLockStatusUpdateVisitor(AbstractVisitor):
    def __init__(self, person1_home, person2_home):
        self.person1_home = person1_home
        self.person2_home = person2_home
    def visit(self, element):
        if self.person1_home:
            element.status = 0
        elif self.person2_home:
            element.status = 1
        else:
            element.status = 1
class CoffeeMachineStatusUpdateVisitor(AbstractVisitor):
    def __init__(self, person1_home, person2_home):
        self.person1_home = person1_home
        self.person2_home = person2_home
    def visit(self, element):
        if self.person1_home:
            if self.person2_home:
                element.status = 2
            else:
                element.status = 3
        elif self.person2_home:
            element.status = 4
        else:
            element.status = 0
class ClockStatusUpdateVisitor(AbstractVisitor):
    def __init__(self, person1_home, person2_home):
        self.person1_home = person1_home
        self.person2_home = person2_home
    def visit(self, element):
        if self.person1_home:
            if self.person2_home:
                pass
            else:
                element.status = '00:01'
        elif self.person2_home:
            element.status = '20:22'
        else:
            pass

class HomeAutomationBootTests(unittest.TestCase):
    def setUp(self):
        self.my_home_system = MyHomeSystem([
            Thermostat("general thermostat"),
            ThemperatureRegulator("thermal regulator"),
            DoorLock("front door lock"),
            CoffeeMachine("coffee machine"),
            Light("Bedroom Light"),
            Clock("system lock")
        ])
    def test_person1_nothome_person2_nothome(self):
        expected_state = map(
            str,
            [
                self.my_home_system.iterable[0].status,
                'off',
                1,
                0,
                0,
                self.my_home_system.iterable[5].status
            ]
        )
        self.visitor = CompositeVisitor(False, False)
        self.my_home_system.accept(self.visitor)
        retrieved_state = sorted([str(x.status) for x in self.my_home_system.iterable])
        self.assertEqual(retrieved_state, sorted(expected_state))
    def test_person1_home_person2_nothome(self):
        pass
    def test_person1_nothome_person2_home(self):
        pass
    def test_person1_home_person2_home(self):
        pass

if __name__ == "__main__":
    unittest.main()
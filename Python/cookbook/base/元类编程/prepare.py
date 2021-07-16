"""
prepare（创建命名空间）-> 依次执行类定义语句 -> new（创建类）-> init（初始化类）
元类定义了prepare以后，会最先执行prepare方法，返回一个空的定制的字典，然后再执行类的语句，
类中定义的各种属性被收集入定制的字典，最后传给new和init方法。
"""
from collections import UserDict


class OrderedClass(type):
    class _member(UserDict):
        def __init__(self):
            super().__init__()
            self.member_names = []

        def __setitem__(self, key, value):
            if key not in self:
                self.member_names.append(key)

            super().__setitem__(key, value)

    @classmethod
    def __prepare__(metacls, cls_name: str, bases: tuple):
        classdict = metacls._member()
        print("prepare return dict id is:", id(classdict))  # 1
        return classdict

    def __new__(metacls, cls_name: str, bases: tuple, classdict: _member):
        print("new get dict id is:", id(classdict))  # 3
        result = type.__new__(metacls, cls_name, bases, dict(classdict))
        result.member_names = classdict.member_names
        print("the class's __dict__ id is:", id(result.__dict__))  # 4
        return result

    def __init__(cls, cls_name: str, bases: tuple, classdict: _member):
        print("init get dict id is ", id(classdict))  # 5
        super().__init__(cls_name, bases, classdict)


class MyClass(metaclass=OrderedClass):
    def method1(self):
        pass

    def method2(self):
        pass

    print("MyClass locals() id is ", id(locals()), locals())  # 2


print("here")
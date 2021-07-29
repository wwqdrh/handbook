# class SingletonObject:

#     class __SingletonObject:
#         def __init__(self):
#             self.val = None
#         def __str__(self):
#             # 这里的 !r 表示使用 repr 输出字符串 repr() 与 str() 的区别，repr的信息更加详细
#             return '{0!r} {1}'.format(self, self.val)

#     instance = None

#     def __new__(cls):
#         if not cls.instance:
#             cls.instance = cls.__SingletonObject()
#         return cls.instance

#     def __getattr__(self, name):
#         return getattr(self.instance, name)

#     def __setattr__(self, name, value):
#         return setattr(self.instance, name, value)

# if __name__ == "__main__":
#     obj1 = SingletonObject()
#     obj2 = SingletonObject()
#     obj1.val = 'object value 1'
#     print('print obj1:', obj1)
#     obj2.val = 'object value 2'
#     print('print obj1:', obj1)
#     print('print obj2:', obj2)


# 元类实现, 调用顺序 type.__init__ type.__call__ type__new__ type__init__
# class SingletonType(type):
#     def __init__(self, *args, **kwargs):
#         print('single_init')
#         super(SingletonType, self).__init__(*args, **kwargs)

#     def __call__(cls, *args, **kwargs):  # 这里的cls，即Foo类
#         print('cls', cls)
#         obj = cls.__new__(cls, *args, **kwargs)
#         cls.__init__(obj, *args, **kwargs)  # Foo.__init__(obj)
#         return obj


# class Foo(metaclass=SingletonType):  # 指定创建Foo的type为SingletonType
#     def __init__(self, name):
#         print('foo_init')
#         self.name = name

#     def __new__(cls, *args, **kwargs):
#         print('foo_new')
#         return object.__new__(cls)


# obj = Foo('xx')

import threading


class SingletonType(type):
    _instance_lock = threading.Lock()

    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            with SingletonType._instance_lock:
                if not hasattr(cls, "_instance"):
                    cls._instance = super().__call__(*args, **kwargs)
        return cls._instance


class Foo(metaclass=SingletonType):
    def __init__(self, name):
        print('here')
        self.name = name


obj1 = Foo('name')
obj2 = Foo('name')
print(obj1, obj2)
print(type(obj1).__dict__, type(obj2).__dict__)

import time

from functools import wraps


def profileing_wrapper(f):
    @wraps(f)
    def wrap_f(*args, **kwargs):
        start_time = time.time()
        result = f(*args, **kwargs)
        end_time = time.time()
        elapesd_time = end_time - start_time
        print("[Time elapsed for n = {}]{}".format(result ,elapesd_time))
        return result
    return wrap_f

def profile_all_class_methods(Cls):
    class ProfiledClass:
        def __init__(self, *args, **kwargs):
            self.inst = Cls(*args, **kwargs)
        # 方法一
        def __getattribute__(self, s):
            # 使用__getattribute__，__getattr__等方法的时候一定要避免写出死循环，不要直接使用self....来获取属性
            # __getattribute__ 是基类object的东西，所以使用的时候可以使用super调用objcet的方法，所以本来是隐式的调用__getattribute__方法的，这里可以显示调用
            x = super().__getattribute__('inst').__getattribute__(s)
            if hasattr(x, '__call__'):
                return profileing_wrapper(x)
            else:
                return x
        # 方法二
        # def __getattribute__(self, s):
        #     # 这里解释一下这种方法也有想要的效果，由于是装饰器必须保留被装饰的原属性，
        #     # 隐式的__getattribute__获取到装饰类的属性后就直接返回了
        #     # 所以这里阻断一下，如果不是获取被装饰类的实例，那么就阻断,如果是的话就返回，虽然没有方法二简洁，但是是一次学习
        #     if s != 'inst':
        #         raise AttributeError
        #     else:
        #         return super().__getattribute__(s)
        # def __getattr__(self, s):
        #     x = getattr(self.inst, s)
        #     if hasattr(x, '__call__'):
        #         return profileing_wrapper(x)
        #     else:
        #         return x
                
    return ProfiledClass

@profile_all_class_methods
class DoMathStuff:
    def __init__(self):
        self.x = 1
        self.y = 2
    def fib(self):
        print(1)

if __name__ == "__main__":
    domathstuff = DoMathStuff()
    print(domathstuff.__dict__)
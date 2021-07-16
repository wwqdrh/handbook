from abc import ABC, abstractmethod

class Target(ABC):
    @abstractmethod
    def request(self):
        print("普通请求")

class Adaptee:
    def specific_request(self):
        print("特殊请求")

class Adapter(Target):
    def __init__(self):
        self.adaptee = Adaptee()

    def request(self):
        self.adaptee.specific_request()

if __name__ == "__main__":
    target = Adapter()
    target.request()
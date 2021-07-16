from weakref import WeakValueDictionary

class EventManager:

    def __init__(self):
        self.__eventMapping = WeakValueDictionary()

    def register(self, message: str, event=None):
        if event is None:
            return lambda event_: self.__eventMapping.update({message: event_}
                                                            ) or event_
        self.__eventMapping[message] = event
        return event

    def unregister(self, message: str):
        self.__eventMapping.pop(message)

    def send_event(self, message: str, *args,
                   **kwargs):
        """
        我们不知道注册的是普通函数还是协程函数
        :param message:
        :param args:
        :param kwargs:
        :return:
        """
        event = self.__eventMapping[message]
        return event(*args, **kwargs)


EVENT = EventManager()


class A:
    def __init__(self):
        EVENT.register('test', self.a)
    # @EVENT.register('test')
    def a(self):
        print("test a")
a = A()
EVENT.send_event('test')

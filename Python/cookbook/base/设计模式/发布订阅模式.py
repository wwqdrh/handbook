'''
与观察者模式类似但又有所不同：观察者模式允许我们将观察的对象解耦出来，从而它们不需要知晓与观察它们
的对象有关的任何信息，但观察者对象仍旧需要它们需要观察哪些对象，这样的处理仍然具有较多的耦合性

发布订阅模式：我们希望观察者或者可被观察对象都不需要知道彼此的信息
'''
class Message:
    def __init__(self):
        self.payload = None
        self.topic = "all"

'''
@class Subscriber 订阅者
'''
class Subscriber:
    def __init__(self, dispatcher, topic):
        dispatcher.subscribe(self, topic)
    def process(self, message):
        print("Message: {}".format(message.payload))

'''
@class Publisher 发布者
'''
class Publisher:
    def __init__(self, dispatcher):
        self.dispatcher = dispatcher
    def publish(self, message):
        self.dispatcher.send(message)

'''
@class Disapther 订阅者与发布者之间的分发器
'''
class Dispatcher:
    def __init__(self):
        self.topic_subscribers = {}
    def subscribe(self, subscriber, topic):
        self.topic_subscribers.setdefault(topic, set()).add(subscriber)
    def unsubscribe(self, subscriber, topic):
        self.topic_subscribers.setdefault(topic, set()).discard(subscriber)
    def unsubscribe_all(self, topic):
        self.subscribers = self.topic_subscribers[topic] = set()
    def send(self, message):
        for subscriber in self.topic_subscribers[message.topic]:
            subscriber.process(message)

if __name__ == "__main__":
    dispatcher = Dispatcher()

    publisher_1 = Publisher(dispatcher)
    subscriber_1 = Subscriber(dispatcher, 'topic1')

    message = Message()
    message.payload = 'My Payload'
    message.topic = 'topic1'

    publisher_1.publish(message)
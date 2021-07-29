import datetime

'''
规则类，用于定义情况
@class Rule
@class Condition
@class And
@class Or
'''
class Rule:
    def __init__(self, conditions, discounts):
        self.conditions = conditions
        self.discounts = discounts
    def evaluate(self, tab):
        if self.conditions.evaluate(tab):
            return self.discounts.calculate(tab)
        return 0

class Conditions:
    def __init__(self, expression):
        self.expression = expression
    def evaluate(self, tab):
        return self.expression.evaluate(tab)

class And:
    def __init__(self, expression1, expression2):
        self.expression1 = expression1
        self.expression2 = expression2
    def evaluate(self, tab):
        return self.expression1.evaluate(tab) and self.expression2.evaluate(tab)

class Or:
    def __init__(self, expression1, expression2):
        self.expression1 = expression1
        self.expression2 = expression2
    def evaluate(self, tab):
        return self.expression1.evaluate(tab) or self.expression2.evaluate(tab)


'''
折扣类，用于定义各种各样的折扣方案
@class PercentageDiscount
@class CheapsFree
'''
class PercentageDiscount:
    def __init__(self, item_type, percentage):
        self.item_type = item_type
        self.percentage = percentage
    def calculate(self, tab):
        return sum(x.cost for x in tab.items if x.item_type == self.item_type) * self.percentage / 100

class CheapsFree:
    def __init__(self, item_type):
        self.item_type = item_type
    def calculate(self, tab):
        try:
            return min(x.cost for x in tab.items if x.item_type == self.item_type)
        except:
            return 0


'''
时间类，用于定义各种各样的时间情况
@class TodayIs
@class TimeIsBetween
@class TodayIsAWeekDay
@class TodayIsAWeekedDay
@class DayOfTheAweek
'''
class TodayIs:
    WeekName = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
    def __init__(self, day_of_week):
        self.day_of_week = day_of_week
    def evaluate(self, tab):
        return self.WeekName[datetime.datetime.today().weekday()] == self.day_of_week.name

class TimeIsBetween:
    def __init__(self, from_time, end_time):
        self.from_time = from_time
        self.end_time = end_time
    def evaluate(self, tab):
        hour_now = datetime.datetime.today().hour
        minute_now = datetime.datetime.today().minute

        from_hour, from_minute = [int(x) for x in self.from_time.split(':')]
        to_hour, to_minute = [int(x) for x in self.end_time.split(':')]

        hour_in_range = from_hour <= hour_now < to_hour
        begin_edge = hour_now == from_hour and minute_now > from_minute
        end_edge = hour_now == to_hour and minute_now < to_minute

        return any([hour_in_range, begin_edge, end_edge])

class TodayIsAWeekDay:
    def __init__(self):
        pass
    def evaluate(self, tab):
        week_days = [
            'Monday',
            'Tuesday',
            'Wednesday',
            'Thursday',
            'Friday'
        ]
        return datetime.datetime.today().weekday() in week_days

class TodayIsAWeekedDay:
    def evaluate(self, tab):
        weekend_days = [
            'Saturday',
            'Sunday'
        ]
        return datetime.datetime.today().weekday() in weekend_days

class DayOfTheWeek:
    def __init__(self, name):
        self.name = name

'''
控制台类
@class Tab
'''
class Tab:
    def __init__(self, customer):
        self.items = []
        self.discounts = []
        self.customer = customer
    def calculate_cost(self):
        return sum(x.cost for x in self.items)
    def calculate_discount(self):
        return sum(x for x in self.discounts)
'''
服务类
@class Item
@class ItemType
@class ItemIsA
@class NumberOfItemOfType 判断所点商品中是否有指定数量
'''
class Item:
    def __init__(self, name, item_type, cost):
        self.name = name
        self.item_type = item_type
        self.cost = cost

class ItemType:
    def __init__(self, name):
        self.name = name

class ItemIsA:
    def __init__(self, item_type):
        self.item_type = item_type
    def evaluate(self, item):
        return self.item_type == item.item_type

class NumberOfItemOfType:
    def __init__(self, number_of_items, item_type):
        self.number = item_type
        self.item_type = number_of_items
    def evaluate(self, tab):
        return len([x for x in tab.items if x.item_type == self.item_type]) == self.number

'''
顾客类，根据可以使用的服务来区分不同的顾客
@class Customer
@class CustomerType 可以使用的服务
@class CustomerIsA
'''
class Customer:
    def __init__(self, customer_type, name):
        self.customer_type = customer_type
        self.name = name

class CustomerType:
    def __init__(self, customer_type):
        self.customer_type = customer_type

class CustomerIsA:
    def __init__(self, customer_type):
        self.customer_type = customer_type
    def evaluate(self, tab):
        return tab.customer.customer_type == self.customer_type

def setup_demo_tab():
    member_customer = Customer(member, 'John')
    tab = Tab(member_customer)

    tab.items.append(Item('Margarita', pizza, 15))
    tab.items.append(Item('Cheddar Melt', burger, 6))
    tab.items.append(Item('Cheddar Melt2', burger, 6))
    tab.items.append(Item('Hawaian', pizza, 12))
    tab.items.append(Item('Latte', drink, 4))
    tab.items.append(Item('Club', pizza, 17))

    return tab

if __name__ == "__main__":
    member = CustomerType('Member')
    pizza = CustomerType('pizza')
    burger = CustomerType('burger')
    drink = CustomerType('drink')
    thursday = DayOfTheWeek('Thursday')

    tab = setup_demo_tab()

    rules = []
    # rules.append(
    #     Rule(
    #         CustomerIsA(member), 
    #         PercentageDiscount('any_item', 15)
    #     )
    # )
    # rules.append(
    #     Rule(
    #         And(TimeIsBetween("17:00", "19:00"), TodayIsAWeekDay()),
    #         PercentageDiscount(drink, 10)
    #     ), 
    # )
    rules.append(
        Rule(
            And(TodayIs(thursday), NumberOfItemOfType(burger, 2)),
            CheapsFree(burger)
        )
    )
    for rule in rules:
        tab.discounts.append(rule.evaluate(tab))
    
    print(
        'Calculated cost: {}\nDiscount applied: {}\n'.format(
            tab.calculate_cost(),
            tab.calculate_discount()
        )
    )
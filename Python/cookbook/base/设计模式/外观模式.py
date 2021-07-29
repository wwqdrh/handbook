'''
外观模式并非只是用于封装单个对象，而是还会提供一个封装器，它会为一组复杂子系统提供一个简化接口，
该接口易于使用，且没有不必要的函数或复杂性
'''

class Invoice:
    def __init__(self, customer):
        pass

class Customer:
    @classmethod
    def fetch(cls, customer_code):
        pass
    def save(self):
        pass

class Item:
    @classmethod
    def fetch(cls, customer_code):
        pass
    def save(self):
        pass

class SaleFacade:
    @staticmethod
    def make_invoice(customer_id):
        return Invoice(Customer.fetch(customer_id))
    @staticmethod
    def make_customer():
        return Customer()
    @staticmethod
    def make_item(item_barcode):
        return Item.fetch(item_barcode)
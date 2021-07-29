class Dispatcher:
    def __init__(self, handlers=[]):
        self.handlers = handlers
    def handle_request(self, request):
        for handler in self.handlers:
            request = handler(request)
        return request

def fun1(instr):
    print(instr)
    return "".join([x for x in instr if x != '1'])

def fun2(instr):
    print(instr)
    return "".join([x for x in instr if x != '2'])

def fun3(instr):
    print(instr)
    return "".join([x for x in instr if x != '3'])

if __name__ == "__main__":
    dispatcher = Dispatcher([
        fun1, fun2, fun3
    ])
    dispatcher.handle_request('12334567843242')
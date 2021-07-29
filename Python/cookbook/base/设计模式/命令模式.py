'''
这个实例将 命令对象 和 接收者简化到了一个函数中
'''

class Invoker:
    def __init__(self):
        self.commands = []
    def add_command(self, command):
        self.commands.append(command)
    def run(self):
        for command in self.commands:
            command['function'](*command['params'])

if __name__ == "__main__":
    def f(string1, string2):
        print("writing {} - {}".format(string1, string2))
    #f = lambda string1, string2: print("writing {} - {}".format(string1, string2))

    invoker = Invoker()
    invoker.add_command({
        'function': f,
        'params': ('command1', 'string1')
    })
    invoker.add_command({
        'function': f,
        'params': ('command2', 'string2')
    })
    invoker.run()
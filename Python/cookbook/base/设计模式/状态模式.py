import curses
import time


'''
定义一系列状态
@abstractClass State 所有状态的基类
@class Standing
@class RunningLeft
@class RunningRight
@class Jumping
@class Crouching

@class StateMachine 状态机类，由上述定义的状态组合
'''
class State:
    def __init__(self, state_macheine):
        self.state_macheine = state_macheine
    def switch(self, in_key):
        if in_key in self.state_macheine.mapping:
            self.state_macheine.state = self.state_macheine.mapping[in_key]
        else:
            self.state_macheine.state = self.state_macheine.mapping['default']

class Standing(State):
    def __str__(self):
        return 'Standing'

class RunningLeft(State):
    def __str__(self):
        return 'RunningLeft'

class RunningRight(State):
    def __str__(self):
        return 'RunningRight'

class Jumping(State):
    def __str__(self):
        return 'Jumping'

class Crouching(State):
    def __str__(self):
        return 'Crouching'

class StateMachine:
    def __init__(self):
        self.standing = Standing(self)
        self.running_left = RunningLeft(self)
        self.running_right = RunningRight(self)
        self.jumping = Jumping(self)
        self.crouching = Crouching(self)
        self.mapping = {
            'a': self.running_left,
            'd': self.running_right,
            's': self.crouching,
            'w': self.jumping,
            'default': self.standing
        }
        self.state = self.standing
    def __str__(self):
        return str(self.state)
    def action(self, in_key):
        self.state.switch(in_key)

if __name__ == "__main__":
    player1 = StateMachine()
    win = curses.initscr()
    curses.noecho()

    win.addstr(0, 0, 'press the key w a s d to initialte actions')
    win.addstr(1, 0, 'press x to quit')
    win.addstr(2, 0, '> ')
    win.move(2, 2)

    while True:
        ch = win.getch()
        if ch is not None:
            win.move(2, 0)
            win.deleteln()
            win.addstr(2, 0, "> ")
            if ch == 120:
                break
            player1.action(chr(ch))
            print(player1.state)
        time.sleep(0.05)
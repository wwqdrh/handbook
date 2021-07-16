"""
管理通过pip安装的三方包，整理信息，依赖，删除的时候将没有被其他包依赖的包一起
Usage pip-tool [command] [package]
command:
    show [package] 按照tree的格式展示出来
    uninstall [package] 将相关依赖一起删除
"""
import argparse
from typing import List
import sys


class ToolException(BaseException):
    def __init__(self, *args, error_code: int):
        super().__init__(*args)
        self.error_code = error_code


class BaseCommand:
    def parse_options(self, extra):
        parser = argparse.ArgumentParser(prog="package", description="package")
        parser.add_argument("package", metavar="package")
        options = vars(parser.parse_args(extra))
        return options


class ShowCommand(BaseCommand):
    """
    显示package的信息
    """
    def __call__(self, package: str):
        """show package的信息"""
        print(package)


class UninstallCommand(BaseCommand):
    """
    删除包的信息
    """
    def __call__(self, package: str):
        """show package的信息"""
        print(package)


def parse_cmdline(args: List[str]):
    parser = argparse.ArgumentParser(prog="pip-tool",
                                     description="pip package manage tool",
                                     usage="pip-tool <command> <package>",
                                     add_help=False)
    parser.add_argument("command",
                        choices=["show", "uninstall"],
                        metavar="command",
                        help="the command to execute (one of: %(choices)s)")
    """
    有时一个脚本可能只解析部分命令行参数，而将其余的参数继续传递给另一个脚本或程序。 
    在这种情况下，parse_known_args() 方法会很有用处。 它的作用方式很类似 parse_args() 
    但区别在于当存在额外参数时它不会产生错误。 而是会返回一个由两个条目构成的元组，
    其中包含带成员的命名空间和剩余参数字符串的列表。
    """
    options, extra = parser.parse_known_args(args)

    command: BaseCommand
    if options.command is None:
        raise ToolException(parser.format_help(), error_code=-10)
    elif options.command == "show":
        command = ShowCommand()
        options = command.parse_options(extra)
        return command, options
    elif options.command == "uninstall":
        command = UninstallCommand()
        options = command.parse_options(extra)
        return command, options


def main():
    try:
        command, options = parse_cmdline(sys.argv[1:])
        command(**options)
        result = 0
    except ToolException as e:
        print(e, file=sys.stdout if e.error_code == 0 else sys.stderr)
        result = e.error_code

    sys.exit(result)


if __name__ == "__main__":
    main()
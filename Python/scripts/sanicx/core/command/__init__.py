"""
定义规范
每一个模块都是一个command，都包含Command类，并且是可调用对象，包入口是向外部暴露相应的公共部分
1、列举出当前包文件下的模块，并且筛选出具有Command的将，文件名与Command的映射关系列出
"""
__all__ = ("CommandUtil", "commands")

import argparse
import collections
import importlib
import os
import pkgutil
import sys
import typing as T

from sanicx.core.exception import CommandError


def _get_commands() -> T.Dict[str, T.Callable]:
    """
    在当前包中访问包含的commands
    """
    res: T.Dict[str, T.Callable] = {}
    path, package = __path__[0], __package__
    for i in filter(
            lambda i_: not i_.ispkg, pkgutil.iter_modules([path])
    ):  # type: pkgutil.ModuleInfo
        command_path = f"{package}.{i.name}"
        command_module = importlib.import_module(command_path)
        command_name = getattr(command_module, "CORE_COMMAND", None)
        if command_name is None:
            continue
        command = getattr(command_module, command_name, None)
        if command is None or not callable(command):
            continue
        res[i.name] = command
    return res


commands: T.Dict[str, T.Callable] = _get_commands()


class _CommandParser(argparse.ArgumentParser):
    """
    Customized ArgumentParser class to improve some error messages and prevent
    SystemExit in several occasions, as SystemExit is unacceptable when a
    command is called programmatically.
    """

    def __init__(
            self, *, missing_args_message=None, called_from_command_line=None, **kwargs
    ):
        self.missing_args_message = missing_args_message
        self.called_from_command_line = called_from_command_line
        super().__init__(**kwargs)

    def parse_args(self, args=None, namespace=None):
        # Catch missing argument for a better error message
        if self.missing_args_message and not (
                args or any(not arg.startswith("-") for arg in args)
        ):
            self.error(self.missing_args_message)
        return super().parse_args(args, namespace)

    def error(self, message):
        if self.called_from_command_line:
            super().error(message)
        else:
            raise CommandError("Error: %s" % message)

    def main_help_text(self, commands_only=False):
        """Return the script's main help text, as a string."""
        if commands_only:
            usage = sorted(commands)
        else:
            usage = [
                "",
                "Type '%s help <subcommand>' for help on a specific subcommand."
                % self.prog,
                "",
                "Available subcommands:",
            ]
            commands_dict = collections.defaultdict(lambda: [])
            for name, _ in commands.items():
                name = name.rpartition(".")[-1]
                commands_dict["sanicx.core"].append(name)
            for app in sorted(commands_dict):
                usage.append("")
                for name in sorted(commands_dict[app]):
                    usage.append("    %s" % name)

        return "\n".join(usage)


class CommandUtil:
    """
    CommandUtil的帮助函数
    """

    parser = _CommandParser(
        usage="%(prog)s subcommand [options] [args]", add_help=False, allow_abbrev=False
    )
    parser.add_argument('command', nargs="?", default="help")
    parser.add_argument("args", nargs="*")  # catch-all

    def __init__(self):
        self.prog_name = os.path.basename(sys.argv[0])
        self.args, self._options = self.parser.parse_known_args(sys.argv[1:])
        if self.prog_name == "__main__.py":
            self.prog_name = "python -m sanicx"
        self.settings_exception = None

    def execute(self):
        """
        cli方法入口
        """
        subcommand = self.args.command
        if subcommand == "help":
            sys.stdout.write(self.parser.main_help_text() + "\n")
        else:
            commands[subcommand](self._options)()

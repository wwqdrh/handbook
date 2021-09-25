"""
@version: 3.9
@title: cli应用程序

- add_argument: 可选参数 action help type nargs
- add_subparsers: 子命令
- add_argument_group: 参数组 只是作为参数的组的概念 实际的解析情况还是与原来相同
- add_mutually_exclusive_group: 互斥
- set_defaults: 设置默认值
- parse_known_args: 只解析已知的部分

ArgumentParser.add_argument_group(title=None, description=None)
在默认情况下，ArgumentParser 会在显示帮助消息时将命令行参数分为“位置参数”和“可选参数”两组。 
当存在比默认更好的参数分组概念时，可以使用 add_argument_group() 方法来创建适当的分组
"""
import argparse


def _parser_builder() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(prog="Test App Prog")

    parser.add_argument("--foo", action="store_true", help="foo help")
    subparsers = parser.add_subparsers(help="子命令 help")
    # create the parser for the "a" command
    parser_a = subparsers.add_parser("a", help="a help")
    parser_a.add_argument("bar", type=int, help="bar help")

    # create the parser for the "b" command
    parser_b = subparsers.add_parser("b", help="b help")
    parser_b.add_argument("--baz", choices="XYZ", help="baz help")

    group = parser.add_argument_group("group")
    group.add_argument("--foo2", help="foo2 help")
    # group.add_argument("bar2", help="bar2 help")

    # parse some argument lists
    print(parser.parse_args(["a", "12"]))
    print(parser.parse_args(["--foo", "b", "--baz", "Z"]))

    return parser


parser = _parser_builder()

if __name__ == "__main__":
    parser.parse_args()

import argparse

from pyup.initialize import command_init
from pyup.generate import command_generate

parser = argparse.ArgumentParser(
    description="pyup, a tookit which speed up your python code by Cython"
)
sub_parser = parser.add_subparsers(help="子命令")


# init
init_parser = sub_parser.add_parser("init", help="初始化项目")
init_parser.add_argument("-f", dest="force", action="store_true", default=False)
init_parser.add_argument("name", nargs=1)
init_parser.set_defaults(func=command_init)

# generator 根据类型提示自动生成pxd代码
gen_parser = sub_parser.add_parser("generator", help="自动生成pxd代码")
gen_parser.add_argument("name", nargs=1)
gen_parser.set_defaults(func=command_generate)

# def compile_():
compile_parser = sub_parser.add_parser("compile", help="编译当前项目")


# def deploy():
deploy_parser = sub_parser.add_parser("deploy", help="保存修改到python path")


def main():
    args = parser.parse_args()
    args.func(args)


if __name__ == "__main__":
    main()

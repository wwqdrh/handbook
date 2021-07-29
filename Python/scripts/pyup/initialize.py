"""
pyup -i [-f] pydantic 项目初始化之类

读取目录
设置配置
"""
import argparse
import pathlib
import os
import os.path
import sys
import subprocess
import shutil
from typing import Optional
import logging


logger = logging.Logger("command-setup")


def _existis_lib(package_name: str, force: bool = False) -> Optional[pathlib.Path]:
    """
    判断所提供的包名，例如pydantic，是否已经安装了
    force = True 表示强制安装，如果不存在则尝试安装
    >>> _existis_lib("rerere") is None
    True
    >>> isinstance(_existis_lib("pip"), pathlib.Path)
    True

    # >>> try:
    # ...     _existis_lib("httpxx", True)
    # ... except subprocess.CalledProcessError as e:
    # ...     e
    # CalledProcessError(1, ['pip', 'install', '--user', 'httpxx'])
    """
    for path in map(lambda i: os.path.join(i, package_name), sys.path):
        if os.path.exists(path):
            return pathlib.Path(path)

    if not force:
        return None

    # check = true 表示如果以非0返回，该进程抛出CalledProcessError, capture_output会捕获输出
    logger.info(f"now execute pip install --user {package_name}, wait !!!")
    subprocess.run(
        ["pip", "install", "--user", package_name], check=True, capture_output=True
    )
    for path in map(lambda i: os.path.join(i, package_name), sys.path):
        if os.path.exists(path):
            return pathlib.Path(path)

    return None


def _copy_lib(package_path: pathlib.Path):
    """
    复制代码到工作目录中
    # >>> _copy_lib(pathlib.Path(_existis_lib("httpx"))) is None
    # True
    """
    target_path = pathlib.Path(os.path.join(os.getcwd(), package_path.name))

    shutil.copytree(package_path, target_path)


def command_init(args: argparse.Namespace):
    """
    Namespace:
        force: bool
        name: list(1)
    """
    package_name, force = args.name[0], args.force
    if (path := _existis_lib(package_name, force)) is None:
        print(f"包名 {package_name} 不存在")
        exit(0)
    _copy_lib(path)


if __name__ == "__main__":
    import doctest

    doctest.testmod()

from __future__ import annotations

import sys
import os
import re
import tempfile
import itertools
import shutil
import argparse
import pathlib


def get_args() -> argparse.Namespace:
    """
    获取命令行参数

    -f ... 需要处理的文件
    """
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument("-f", dest="files", nargs="*")
    arg_parser.add_argument("-uf", dest="ufiles", nargs="*")

    return arg_parser.parse_args(sys.argv[1:])


def action(file: pathlib.Path, types: str = "add") -> None:
    """
    处理给定的路径下文件的处理， 为每一个节点中的class加上tw前缀
    """

    def _add_prefix(line: str) -> str:
        if not (tw_match := re.search(r"class=[\"\'](.*)[\"\']", line)):
            return line

        tw_class = tw_match.group(1)
        line = re.sub(r"class=[\"\'].*[\"\']", "{tw_class}", line)
        return line.format(
            tw_class='class="{}"'.format(
                " ".join(map(lambda i: f"tw-{i}", tw_class.split(" ")))
            )
        )

    def _remove_prefix(line: str) -> str:
        if not (tw_match := re.search(r"class=[\"\'](.*)[\"\']", line)):
            return line

        tw_class = tw_match.group(1)
        line = re.sub(r"class=[\"\'].*[\"\']", "{tw_class}", line)
        return line.format(
            tw_class='class="{}"'.format(
                " ".join(map(lambda i: i.lstrip("tw-"), tw_class.split(" ")))
            )
        )

    with open(file, "r", encoding="utf8") as source_file, open(
        (temp := tempfile.mkstemp()[1]), "w", encoding="utf8"
    ) as temp_file:
        while (line := source_file.readline()) != "":
            if types == "add":
                temp_file.write(_add_prefix(line))
            elif types == "remove":
                temp_file.write(_remove_prefix(line))
    os.remove(file)
    shutil.move(temp, file)


if __name__ == "__main__":
    args = get_args()

    if args.files is not None:
        for file in map(pathlib.Path, args.files):
            if not os.path.exists(file):
                print("{}不存在".format(str(file)))
                continue
            action(file)
    
    if args.ufiles is not None:
        for file in map(pathlib.Path, args.ufiles):
            if not os.path.exists(file):
                print("{}不存在".format(str(file)))
                continue
            action(file, types="remove")

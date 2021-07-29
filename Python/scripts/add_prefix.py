"""
cli应用程序，为文件中的内容某个特定的class添加或者删除前缀
"""
import sys
import os
import re
import tempfile
import shutil
import argparse
import pathlib


arg_parser = argparse.ArgumentParser()
arg_parser.add_argument("-f", dest="files", nargs="*")
arg_parser.add_argument("-uf", dest="ufiles", nargs="*")


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


def main():
    args = arg_parser.parse_args(sys.argv[1:])

    if args.files is not None:
        for file in map(pathlib.Path, args.files):
            if not os.path.exists(file):
                print("{}不存在".format(str(file)))
                continue
            action(file)
    elif args.ufiles is not None:
        for file in map(pathlib.Path, args.ufiles):
            if not os.path.exists(file):
                print("{}不存在".format(str(file)))
                continue
            action(file, types="remove")
    else:
        arg_parser.print_help()


if __name__ == "__main__":
    main()

from configparser import ConfigParser
import os
from typing import Tuple, Generator, TypedDict, List


class FileInfo(TypedDict):
    fullpath: str    # 全路径名
    name: str    # 文件名
    path: str    # 相对于工作目录的路径
    level: int    # 相对于工作目录的层级
    isFile: bool    # 是否是文件


class DirectoryManage:
    """
    文件系统目录管理, 递归给定文件夹的信息
    """

    def __init__(self, cwd_dir: str = None):
        self.cwd_dir = os.getcwd()    # 当前的工作目录

    def gen_file(
        self, folder: "str",
        exclude: Tuple = (".",)) -> Generator[FileInfo, None, None]:
        # 给定文件夹的名字生成当前文件夹下的文件信息, 默认以.开头的文件夹不进行遍历
        pathStack: List[FileInfo] = [{
            "fullpath": os.path.join(self.cwd_dir, folder),
            "name": folder,
            "path": "",
            "level": -1,
            "isFile": False
        }]    # 目录树的栈结构，用来迭代出树的结构
        while pathStack:
            curFile = pathStack.pop()
            if os.path.isdir(curFile["fullpath"]):
                # 当是一个文件夹的时候将里面的元素遍历出来添加到栈中，然后把当前的文件夹信息返回出去
                for child in os.listdir(curFile["fullpath"]):
                    if any(child.startswith(exc) for exc in exclude):
                        continue
                    childInfo: FileInfo = {
                        "fullpath": os.path.join(curFile["fullpath"], child),
                        "name": child,
                        "path": "{}/{}".format(curFile["path"], child),
                        "level": curFile["level"] + 1,
                        "isFile": False
                    }
                    if os.path.isfile(
                            childInfo["fullpath"]):    # 如果是文件将isFile设为True
                        childInfo["isFile"] = True
                    pathStack.append(childInfo)
            yield curFile


class IniManage:

    def __init__(self, ini_file: str):
        self.iniFile = ini_file
        self.config = ConfigParser()
        self.config.read(ini_file)

    def write_config(self):
        with open(self.iniFile, mode="w") as f:
            self.config.write(f)

    def get_config(self, name: str, section: str = "PATH") -> str:
        return self.config[section][name]

    def has_config(self, name: str, section: str = "PATH") -> bool:
        return name in self.config[section]

    def update_ini(self, name: str, docs: str = None):
        """
        修改配置文件
        """
        if name and docs:    # 如果两个都有那么就是新增或者修改
            self.config["PATH"][name] = docs
        elif name:    # 只有name，docs为None，说明需要删除信息
            if not self.has_config(name):
                raise Exception(f"{name} docs 不存在")
            del self.config["PATH"][name]

        self.write_config()


class MarkdownManage:    # 解析markdown语法
    li_format = staticmethod(lambda word, level: "{}* {}".format(level*'\t', word))
    url_format = staticmethod(lambda word, url: "[{}]({})".format(word, url))


if __name__ == "__main__":
    dire = DirectoryManage()
    prefix = "--"
    for i in dire.gen_file(".", exclude=(".", "__pycache__")):
        print("{}: {}".format(prefix * i["level"], i["name"]))

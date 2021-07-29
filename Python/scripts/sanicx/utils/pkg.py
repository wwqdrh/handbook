import os
import re
import shutil
import zipfile
from importlib import import_module

from sanicx.core.exception import FileExtensionException


def lazy_load(module: str, attr: str):
    def _lazy(*args, **kwargs):
        mod = import_module(module)
        return getattr(mod, attr)(*args, **kwargs)

    return _lazy


def unzip(path: str, target_path: str = ".", delete: bool = False):
    """
    解压zip文件
    :param delete:
    :param path: 全限定路径
    :param target_path:
    :return:
    """
    zip_file = None
    try:
        if target_path == ".":
            target_path = re.match(r"(.*)\.zip$", path).group(1)
        zip_file = zipfile.ZipFile(path)
        zip_list = zip_file.namelist()  # 得到压缩包里所有文件
        for f in zip_list:
            zip_file.extract(f, target_path)  # 循环解压文件到指定目录
    except AttributeError as e:
        raise FileExtensionException("只支持zip压缩文件") from e
    else:
        return target_path
    finally:
        zip_file is not None and zip_file.close()
        delete and os.remove(path)


def search_entry(
    path: str, entry_file: str, template_path: str = ".", delete: bool = False
):
    """
    解压完成的文件夹可能是存在嵌套路径的，因此定义一个入口
    找到入口所在的文件夹路径，将这里面的文件全部移动到path路径下
    :param delete:
    :param template_path:
    :param path:
    :param entry_file:
    :return:
    """
    temp_path_name = f"{path}_"

    try:
        os.rename(path, temp_path_name)
        target_path = next(
            root for root, dirs, files in os.walk(temp_path_name) if entry_file in files
        )
        if template_path == ".":
            if target_path == temp_path_name:  # 就在第一层
                os.rename(temp_path_name, path)
            else:
                shutil.move(target_path, path)
                shutil.rmtree(temp_path_name)
        else:
            # 指定了模板的目录
            root, contents = next(
                (root, [*dirs, *files]) for root, dirs, files in os.walk(target_path)
            )
            for content in map(lambda i: os.path.join(root, i), contents):
                shutil.move(content, template_path)
            os.rename(temp_path_name, path)
    except StopIteration as e:
        os.rename(temp_path_name, path)
        raise FileExtensionException(f"该文件夹下不包含该入口文件[{entry_file}]") from e
    finally:
        delete and shutil.rmtree(path)

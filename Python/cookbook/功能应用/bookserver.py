from http.server import HTTPServer, SimpleHTTPRequestHandler
import argparse
import sys
import os
from typing import ClassVar, List
from utils.native.file import IniManage, DirectoryManage, MarkdownManage
from functools import partial


class GitServer(HTTPServer):
    """
    server对象
    """
    allow_reuse_address = True  # 使得服务socket能够重用

    class __Handler(SimpleHTTPRequestHandler):
        def send_head(self):
            """
            重写send_head，加上自己的标记头, 使得每一个请求存在有效期
            """
            import os
            import urllib
            import email
            import datetime
            from http import HTTPStatus

            path = self.translate_path(self.path)
            f = None
            if os.path.isdir(path):
                parts = urllib.parse.urlsplit(self.path)
                if not parts.path.endswith('/'):
                    # redirect browser - doing basically what apache does
                    self.send_response(HTTPStatus.MOVED_PERMANENTLY)
                    new_parts = (parts[0], parts[1], parts[2] + '/',
                                    parts[3], parts[4])
                    new_url = urllib.parse.urlunsplit(new_parts)
                    self.send_header("Location", new_url)
                    self.end_headers()
                    return None
                for index in "index.html", "index.htm":
                    index = os.path.join(path, index)
                    if os.path.exists(index):
                        path = index
                        break
                else:
                    return self.list_directory(path)
            ctype = self.guess_type(path)
            # check for trailing "/" which should return 404. See Issue17324
            # The test for this was added in test_httpserver.py
            # However, some OS platforms accept a trailingSlash as a filename
            # See discussion on python-dev and Issue34711 regarding
            # parseing and rejection of filenames with a trailing slash
            if path.endswith("/"):
                self.send_error(HTTPStatus.NOT_FOUND, "File not found")
                return None
            try:
                f = open(path, 'rb')
            except OSError:
                self.send_error(HTTPStatus.NOT_FOUND, "File not found")
                return None

            try:
                fs = os.fstat(f.fileno())
                # Use browser cache if possible
                if ("If-Modified-Since" in self.headers
                        and "If-None-Match" not in self.headers):
                    # compare If-Modified-Since and time of last file modification
                    try:
                        ims = email.utils.parsedate_to_datetime(
                            self.headers["If-Modified-Since"])
                    except (TypeError, IndexError, OverflowError, ValueError):
                        # ignore ill-formed values
                        pass
                    else:
                        if ims.tzinfo is None:
                            # obsolete format with no timezone, cf.
                            # https://tools.ietf.org/html/rfc7231#section-7.1.1.1
                            ims = ims.replace(tzinfo=datetime.timezone.utc)
                        if ims.tzinfo is datetime.timezone.utc:
                            # compare to UTC datetime of last modification
                            last_modif = datetime.datetime.fromtimestamp(
                                fs.st_mtime, datetime.timezone.utc)
                            # remove microseconds, like in If-Modified-Since
                            last_modif = last_modif.replace(microsecond=0)

                            if last_modif <= ims:
                                self.send_response(HTTPStatus.NOT_MODIFIED)
                                self.end_headers()
                                f.close()
                                return None

                self.send_response(HTTPStatus.OK)
                self.send_header("Content-type", ctype)
                self.send_header("Content-Length", str(fs[6]))
                self.send_header("Last-Modified",
                                    self.date_time_string(fs.st_mtime))
                self.send_header("Cache-Control", "max-age=60")  # 添加缓存响应头
                self.end_headers()
                return f
            except:
                f.close()
                raise

    def __init__(self,
                 host: str,
                 port: int,
                 directory: str):
        super().__init__((host, port), partial(self.__Handler, directory=directory))
        self.directory = directory  # 保存一下当前服务运行的目录
    
    def serve_forever(self):
        import subprocess
        cur_cwd = os.getcwd()
        os.chdir(self.directory)
        print("服务开启")
        subprocess.run(["open", "http://localhost:9999"])
        super().serve_forever()
        os.chdir(cur_cwd)

class BookApp:
    work_dir: ClassVar[str]
    docs_dir: ClassVar[str]
    docs_ini: ClassVar[str]
    cwd_dir: ClassVar[str]

    argument = argparse.ArgumentParser(description="my own web docs server")
    argument.add_argument("-a", dest="add", nargs=2, help="新增一个docs网站", metavar=("name", "docsdir"))
    argument.add_argument("-d", dest="dele", type=str, help="删除给定名字的docs", metavar="delename")
    argument.add_argument("-s", dest="server", type=str, help="启动服务")
    argument.add_argument("--list", dest="list", action="store_true")
    argument.add_argument("--parse", dest="parse", metavar="target folder", help="解析指定文件夹内的内容生成目录")

    def __new__(cls, *args, **kwargs):
        cls._check_runtime()  # 环境检查, 设置相应的配置变量
        # cls._check_config()  # 处理配置变量
        return super().__new__(cls, *args, **kwargs)
    
    def __init__(self):
        self.iniManage = IniManage(self.docs_ini)
        self.directManage = DirectoryManage(self.cwd_dir)
        self.server_factory = GitServer
        
        self._check_config()
    
    @classmethod
    def _check_runtime(cls):
        """
        检查当前运行环境
        """
        import sys
        import os
        from pathlib import Path

        WorkDir: str = Path(__file__).parents[0]
        Docs = os.path.join(WorkDir, "docs")
        DocsIni = os.path.join(Docs, "docs.ini")

        if not os.path.exists(Docs): os.mkdir(Docs)
        if not os.path.exists(DocsIni):
            with open(DocsIni, mode="w"):
                pass

        cls.work_dir = WorkDir
        cls.docs_dir = Docs
        cls.docs_ini = DocsIni
        cls.cwd_dir = os.getcwd()
    
    def _check_config(self):
        """
        配置变量
        """
        config = self.iniManage.config
        if "SERVER" not in config:
            config["SERVER"] = {
                "HOST": "127.0.0.1",
                "PORT": "9999"
            }
            self.iniManage.write_config()
        if "PATH" not in config: config["PATH"] = {}

    def from_cli(self, arg_list: List[str], debug: bool = False):
        """
        传入命令行参数然后运行
        """
        try:
            if not arg_list: 
                self.argument.print_help()
                return
            args = self.argument.parse_args(arg_list)
            if argument := args.add:
                self.add_docs(argument[0], argument[1])
            elif argument := args.dele:
                self.del_docs(argument)
            elif argument := args.server:
                self.start_server(argument)
            elif argument := args.list:
                self.list_all()
            elif argument := args.parse:
                self.parse(argument)
        except Exception as e:
            if debug: raise e
            self.argument.print_help()
            print(e)
        except KeyboardInterrupt:
            print("程序结束")
    
    def add_docs(self, name: str, dirname: str):
        target_docs_dir = os.path.join(self.cwd_dir, dirname)
        target_entry_file = os.path.join(target_docs_dir, "index.html")
        new_docs = f"{name}-{dirname}"
        new_docs_path = os.path.join(self.docs_dir, new_docs)
        if not os.path.exists(target_docs_dir) \
                or not os.path.exists(target_entry_file):
            raise Exception("目标文件夹不存在或者文件夹中不包含入口index.html")

        import shutil
        shutil.copytree(target_docs_dir, new_docs_path)
        self.iniManage.update_ini(name, new_docs)
    
    def del_docs(self, name: str):
        self.iniManage.update_ini(name)
    
    def start_server(self, name: str):
        """
        开启服务
        """
        if not self.iniManage.has_config(name): raise Exception("当前docs不存在无法启动服务")

        host = self.iniManage.get_config("host", "SERVER")
        port = int(self.iniManage.get_config("port", "SERVER"))
        directory = self.iniManage.get_config(name)
        directory_path = os.path.join(self.docs_dir, directory)
        self.server_factory(host, port, directory_path).serve_forever()

    def list_all(self):
        """
        展示当前服务器所有的docs
        """
        if books := self.iniManage.config["PATH"].keys():
            print("\n".join(books))
        else:
            print("无")
    
    def parse(self, folder: str):
        """
        传入当前工作目录下的一个文件夹的名字，将其中的文件结构转换为markdown格式生成目录信息
        """
        lines = ["# 目录\n"]
        for line in self.directManage.gen_file(folder, exclude=(".", "summary.md")):
            if line["level"] == -1: continue  # 根节点
            url = line["path"] if line[
                "isFile"] else f"{line['path']}/README.md"
            name, ext = os.path.splitext(line["name"])
            line_str = MarkdownManage.li_format(
                MarkdownManage.url_format(name, url),
                                        line["level"])
            lines.append(line_str + "\n")
        
        with open(os.path.join(folder, "summary.md"), mode="wt") as f:
            f.writelines(lines)

def main():
    bookApp = BookApp()
    bookApp.from_cli(sys.argv[1:])

if __name__ == "__main__":
    main()
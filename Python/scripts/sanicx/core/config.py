"""
TODO 定义sanicx的核心配置项
以及能够读取应用的配置内容
"""


class Base:
    __conf__ = "base"

    TIME_ZONE = "Asia/Shanghai"


class Server:
    __conf__ = "server"

    port: int = 8080


class Db:
    __conf__ = "db"

    driver: str = "mysql+mysqlconnector"

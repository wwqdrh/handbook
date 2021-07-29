"""
etcd库的处理
"""
from typing import (
    NamedTuple,
    Union,
    TypeVar,
    List,
    Callable,
    Optional,
    Dict,
    List,
    Type,
)
from functools import wraps
from weakref import CallableProxyType, proxy, WeakSet
from threading import Thread
from concurrent.futures import ThreadPoolExecutor, Future

from etcd3 import Etcd3Client

from beopcommon.error import InputDataError

# [stub]
server_key = str
etcd_key = str
cache_key = str

_conn_cache: Dict[str, Etcd3Client] = {}  # type: ignore


class ClientConf(NamedTuple):
    host: str = "localhost"
    port: int = 3306
    user: str = "admin"
    password: str = "123456"

    @property
    def conn_str(self) -> str:
        return f"{self.host}:{self.port}"

    def conn(self) -> Union[Etcd3Client, Exception]:
        try:
            conn_str = self.conn_str
            if conn_str in _conn_cache:
                conn = self.__conn_cache[conn_str]
            else:
                conn = Etcd3Client(
                    host=self.host,
                    port=self.port,
                    user=self.user,
                    password=self.password,
                )
                _conn_cache[conn_str] = conn
        except Exception:
            conn = Exception("连接失败")
        finally:
            return conn


class ConfigField:
    def __init__(self, key: str, value: str = None):
        self.key: str = key
        self.value: Optional[str] = value

    def __str__(self):
        return self.value


class EtcdConfig:
    env_name: str = "_default"
    conn: CallableProxyType  # Etcd3Client
    data_access: Optional["DataAccess"] = None

    def get(self, key: str) -> ConfigField:
        if not self.data_access:
            return ConfigField(key=key)

        return self.data_access.get(key, self.env_name, self.conn)


# [helper]
def daemon_thread(fn: Callable) -> Callable[..., Thread]:
    @wraps(fn)
    def _wrap(*args, **kwargs) -> Thread:
        return Thread(target=fn, args=args, kwargs=kwargs, daemon=True)

    return _wrap


# [secondry]
class DataAccess:
    def __new__(cls, *args, **kwargs):
        try:
            instance = cls.__instance
        except AttributeError:
            instance = super().__new__(cls)
            cls.__instance = instance
        finally:
            return instance

    def __init__(self, domain: str):
        self.__root_prefix = f"{domain}.app"  # prod.app
        self.__local_cache: Dict[cache_key, ConfigField] = dict()
        self.__listener_key: Dict[server_key, Dict[etcd_key, WeakSet]] = {}

    @daemon_thread
    def conn_watch_prefix(self, conn: CallableProxyType):
        """
        conn: Etcd3Client的弱引用
        """
        conn_str: server_key = conn._url
        events_iterator, cancel = conn.watch_prefix(self.__root_prefix)
        for event in map(str, events_iterator):
            # <class 'etcd3.events.PutEvent'> key=b'/testkey/5' value=b'1234'
            key = event.split(" ")[2].split("=")[1][2:-1]
            value = event.split(" ")[3].split("=")[1][2:-1]
            self.notify_listener(conn_str, key, value)

    def wrap_key(self, key: str, env_name: str) -> etcd_key:
        return f"{self.__root_prefix}.{env_name}.{key}"

    def unwrap_key(self, key: etcd_key, env_name: str) -> str:
        return key.replace(f"{self.__root_prefix}.{env_name}.", "")

    def wrap_cache_key(self, key: etcd_key, conn_str: server_key) -> cache_key:
        return f"{conn_str}@{key}"

    def notify_listener(self, server: server_key, key: str, value: str):
        try:
            listener = self.__listener_key[server]
            for config in listener[key]:  # type: ConfigField
                config.value = value
        except KeyError:
            pass

    def register_listener(self, config: "EtcdConfig"):
        conn_str: server_key = config.conn._url

        if conn_str not in self.__listener_key:  # 开启监听
            self.conn_watch_prefix(config.conn).start()

        listener = self.__listener_key.setdefault(conn_str, {})
        for field in (
            val
            for key, val in config.__dict__.items()
            if key.isupper() and isinstance(val, ConfigField)
        ):  # type: ConfigField
            key = self.wrap_key(field.key, config.env_name)
            listener.setdefault(key, WeakSet()).add(field)

    def __get(self, key: str, env_name: str, conn: Etcd3Client) -> ConfigField:
        if env_name == "_default":  # 顶级
            default_config = self.wrap_key(key, "_default")
            result = conn.get(default_config)[0]
            # result = self.client.get(f"{self.defautl_config_prefix}{key}")[0]
            return ConfigField(key=key, value=result and result.decode("utf8"))
        env_prefix = f"{self.__root_prefix}.{env_name}"
        result = conn.get(f"{env_prefix}.{key}")[0]
        if result is None:
            parent_prefix = f"{env_prefix}._extends"
            parent_env = conn.get(parent_prefix)[0]
            parent_env = (
                parent_env and parent_env.decode("utf8")[1:-1]
            ) or "_default"  # 存在parentenv就解码 否则 默认为~
            return self.__get(key, parent_env, conn)
        else:
            return ConfigField(key=key, value=result and result.decode("utf8"))

    def get(self, key: str, env_name: str, conn: Etcd3Client) -> ConfigField:
        # conn_str = conn._url
        # if key.startswith("/"):  # 说明是环境
        #     result = conn.get(f"{self.__root_prefix}{key}")[0]
        #     return ConfigField(key=key, value=result and result.decode("utf8"))
        # else:
        conn_str = conn._url
        etcd_key = self.wrap_key(key, env_name)
        cache_key = self.wrap_cache_key(etcd_key, conn_str)
        if cache_key not in self.__local_cache:
            field = self.__get(key, env_name=env_name, conn=conn)
            if field.value is None:
                return field
            self.__local_cache[cache_key] = field
        return self.__local_cache[cache_key]

    def put(self, key: str, value: str, env_name: str, conn: Etcd3Client):
        """
        更新本地缓存以及etcd服务
        # TODO 暂时不支持这里更新环境
        """
        key = self.wrap_key(key, env_name)
        cache = self.wrap_cache_key(key, conn._url)
        if cache in self.__local_cache:
            self.__local_cache[cache].value = value
        return conn.put(key, value)

    def delete(self, key: str, env_name: str, conn: Etcd3Client):
        # if key.startswith("/"):
        #     raise InputDataError("暂不支持删除环境")
        etcd_key = self.wrap_key(key, env_name)
        cache = self.wrap_cache_key(etcd_key, conn._url)
        if cache in self.__local_cache:
            del self.__local_cache[cache]
        return conn.delete(etcd_key)


# [main]
class EtcdCluster:
    def __init__(
        self, conn: Etcd3Client, domain: str, env_name: str, data_access: DataAccess
    ):
        self.__domain = domain  # prod, dev, 一个进程只有一种domain
        self.env_name = env_name  # cn1, us2
        self.conn = conn
        self.data_access: DataAccess = data_access

    @property
    def members(self):
        return self.conn.members

    @classmethod
    def factory(
        cls, configs: List[ClientConf], domain: str, env_name: str = "_default"
    ) -> "EtcdCluster":
        try:
            conn = next(
                conn
                for conn in map(lambda i: i.conn(), configs)
                if not isinstance(conn, Exception)
            )
        except StopIteration as e:
            raise InputDataError("config is invalid, connection is Error")
        else:
            return cls(
                conn=conn,
                domain=domain,
                env_name=env_name,
                data_access=DataAccess(domain=domain),
            )

    def set_env(self, env_name: str):
        self.env_name = env_name

    def config(self, env_name: str = None) -> Callable:
        def _config(conf: Type[EtcdConfig]):
            @wraps(conf)
            def _wrap(*args, **kwargs) -> EtcdConfig:
                conf.env_name = env_name or self.env_name
                conf.conn = proxy(self.conn)
                conf.data_access = self.data_access
                ins = conf(*args, **kwargs)  # type: ignore
                self.data_access.register_listener(ins)
                return ins

            return _wrap

        return _config

    def put(self, key: str, value: str):
        self.data_access.put(key, value, self.env_name, self.conn)

    def get(self, key: str) -> ConfigField:
        return self.data_access.get(key, self.env_name, self.conn)

    def delete(self, key: str):
        return self.data_access.delete(key, self.env_name, self.conn)

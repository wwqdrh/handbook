"""
etcd库的处理
"""
import ast
from typing import (
    NamedTuple,
    Callable,
    Optional,
    Dict,
    List,
    Type,
    Any,
    Tuple,
    Set,
    Generator,
)
from functools import wraps
from weakref import CallableProxyType, proxy, WeakSet
from threading import Thread
from etcd3 import Etcd3Client
from beopcommon.error import InputDataError

# [stub]
ServerKey = str
EtcdKey = str
CacheKey = str
_conn_cache: Dict[str, Etcd3Client] = {}  # type: ignore


def daemon_thread(fn: Callable) -> Callable[..., Thread]:
    @wraps(fn)
    def _wrap(*args, **kwargs) -> Thread:
        return Thread(target=fn, args=args, kwargs=kwargs, daemon=True)

    return _wrap


class ClientConf(NamedTuple):
    host: str = "localhost"
    port: int = 3306
    user: str = "admin"
    password: str = "123456"

    @property
    def conn_str(self) -> str:
        return f"{self.host}:{self.port}"

    def conn(self) -> Etcd3Client:
        conn_str = self.conn_str
        if conn_str in _conn_cache:
            conn = _conn_cache[conn_str]
        else:
            conn = Etcd3Client(
                host=self.host,
                port=self.port,
                user=self.user,
                password=self.password,
            )
            _conn_cache[conn_str] = conn
        return conn


class ConfigField:
    def __init__(self, *key: str, value: List[Any] = None, prop: Callable = None):
        self.key: Tuple[str, ...] = key
        self._val: Optional[List[str]] = value
        self._prop: Optional[Callable] = prop

    @property
    def value(self):
        val = self._val[0] if len(self.key) == 1 else self._val
        if self._prop is None:
            return val
        else:
            return self._prop(*self._val)

    def set_value(self, key, value):
        for idx, val in enumerate(self.key):
            if val == key:
                self._val[idx] = value

    def get_value(self, key):
        idx = self.key.index(key)
        return self._val[idx]


class EtcdConfig:
    env_name: str = "_default"
    conn: CallableProxyType  # Etcd3Client
    data_access: Optional["DataAccess"] = None

    def __getattribute__(self, item):
        if item.startswith("#"):  # 获取配置项的原始节点
            res = super().__getattribute__(item[1:])
        else:
            res = super().__getattribute__(item)
            if item.upper() == item:
                # guess config
                try:
                    res = res.value
                except AttributeError:
                    pass
        return res

    def get(self, *key: str, prop: Callable = None) -> ConfigField:
        if not self.data_access:
            return ConfigField(*key, prop=prop)
        return self.data_access.get(
            *key, env_name=self.env_name, conn=self.conn, prop=prop
        )


class DataAccess:
    def __new__(cls, *args, **kwargs):
        try:
            instance = cls.__instance
        except AttributeError:
            instance = super().__new__(cls)
            cls.__instance = instance
        return instance

    def __init__(self, domain: str):
        self.__root_prefix = f"{domain}.app"  # prod.app
        self.__local_cache: Dict[CacheKey, List[ConfigField]] = dict()
        self.__listener_key: Dict[ServerKey, Dict[EtcdKey, WeakSet]] = {}

    @daemon_thread
    def conn_watch_prefix(self, conn: CallableProxyType):
        """
        conn: Etcd3Client的弱引用
        """
        # noinspection PyProtectedMember
        conn_str: ServerKey = conn._url
        events_iterator, cancel = conn.watch_prefix(self.__root_prefix)
        for event in map(str, events_iterator):
            # <class 'etcd3.events.PutEvent'> key=b'/testkey/5' value=b'1234'
            key = event.split(" ")[2].split("=")[1][2:-1]
            value = event.split(" ")[3].split("=")[1][2:-1]
            self.notify_listener(conn_str, key, value)

    def wrap_key(self, key: str, env_name: str) -> EtcdKey:
        return f"{self.__root_prefix}.{env_name}.{key}"

    @classmethod
    def wrap_cache_key(cls, key: EtcdKey, server_key: str) -> CacheKey:
        return f"{server_key}@{key}"

    def get_parent_prefix(self, env_name: str, conn: Etcd3Client) -> str:
        env_prefix = f"{self.__root_prefix}.{env_name}"
        parent_prefix = f"{env_prefix}._extends"
        parent_env = conn.get(parent_prefix)[0]
        parent_env = (
            parent_env and parent_env.decode("utf8")[1:-1]
        ) or "_default"  # 存在parentenv就解码 否则 默认为~
        return parent_env

    def notify_listener(self, server_key: str, key: str, value: str):
        # 还需要向父级的配置更新 prod.app.sys.log->prod.app.sys->prod.app
        def _parent_iter(key) -> Generator:
            key_list = key.split(".")
            for idx in range(len(key_list), 0, -1):
                yield ".".join(key_list[:idx])

        listener = self.__listener_key[server_key]
        unwrap_key = ".".join(key.split(".")[3:])
        value = ast.literal_eval(value)
        for cur_key in _parent_iter(key):
            try:
                for config in listener[cur_key]:  # type: ConfigField
                    config.set_value(unwrap_key, value)
            except KeyError:
                pass
            except BaseException as e:
                print(e)

    def register_listener(self, config: "EtcdConfig"):
        # noinspection PyProtectedMember
        conn_str: ServerKey = config.conn._url
        if conn_str not in self.__listener_key:  # 开启监听
            self.conn_watch_prefix(config.conn).start()
        listener = self.__listener_key.setdefault(conn_str, {})
        conn = config.conn
        for field in (
            val
            for key, val in config.__dict__.items()
            if key.isupper() and isinstance(val, ConfigField)
        ):  # type: ConfigField
            for key in field.key:
                env_name = config.env_name
                for _ in range(5):  # 最大层级
                    key_str = self.wrap_key(key, env_name)
                    listener.setdefault(key_str, WeakSet()).add(field)
                    if env_name == "_default":
                        break
                    else:
                        env_name = self.get_parent_prefix(env_name, conn)

    def __get(self, key: str, env_name: str, conn: Etcd3Client) -> bytes:
        if env_name == "_default":  # 顶级
            default_config = self.wrap_key(key, "_default")
            return conn.get(default_config)[0]
            # return ConfigField(key=key, value=result and result.decode("utf8"))
        env_prefix = f"{self.__root_prefix}.{env_name}"
        result = conn.get(f"{env_prefix}.{key}")[0]
        if result is None:
            parent_env = self.get_parent_prefix(env_name, conn)
            result = self.__get(key, parent_env, conn)
        return result

    def __get_dict(
        self, prefix: str, env_name: str, conn: Etcd3Client
    ) -> Optional[dict]:
        # 需要不断往上找找到所有父级配置，添加缺少的
        result: Dict[str, Any] = {}
        prefix_len = prefix.count(".") + 4

        def _get(env_name_: str):
            server_prefix = self.wrap_key(prefix, env_name_)
            cfgs = conn.get_prefix(server_prefix)
            cfg_tree: Dict[str, Any] = {}
            for value, meta in cfgs:
                key = meta.key.decode()
                typed_value = ast.literal_eval(value.decode())
                key_parts = key.split(".")
                if len(key_parts) < 4:
                    continue
                right = len(key_parts)
                cfg_node = result
                for key_part in key_parts[
                    prefix_len:right
                ]:  # prod.app._default.i18n....
                    cfg_node = cfg_node.setdefault(key_part, {})
                last = ".".join(key_parts[right : len(key_parts)])
                cfg_node.setdefault(last, typed_value)

        for _ in range(5):  # 最大层级
            _get(env_name)
            if env_name == "_default":
                break
            else:
                env_name = self.get_parent_prefix(env_name, conn)
        return result or None

    def get(
        self, *key: str, env_name: str, conn: Etcd3Client, prop: Callable = None
    ) -> ConfigField:
        conn_str = conn._url
        cache_keys, values = set(), []
        for item_key in key:
            etcd_key = self.wrap_key(item_key, env_name)
            cache_key = self.wrap_cache_key(etcd_key, conn_str)
            if cache_key not in self.__local_cache:
                origin_data = self.__get(
                    item_key, env_name=env_name, conn=conn
                )  # 特定key
                if origin_data is not None:
                    value = origin_data and ast.literal_eval(origin_data.decode())
                else:
                    # TODO 当key只是一个前缀的时候需要把数据获取出来作为dict
                    value = self.__get_dict(item_key, env_name=env_name, conn=conn)
            else:
                value = self.__local_cache[cache_key][0].get_value(item_key)
            cache_keys.add(cache_key)
            values.append(value)
        field = ConfigField(*key, value=values, prop=prop)
        if field.value is None:
            return field
        for cache_key in cache_keys:
            self.__local_cache.setdefault(cache_key, []).append(field)
        return field

    def put(self, key: str, value: Any, env_name: str, conn: Etcd3Client):
        """
        更新本地缓存以及etcd服务
        # TODO 暂时不支持这里更新环境
        """
        etcd_key = self.wrap_key(key, env_name)
        # noinspection PyProtectedMember
        cache = self.wrap_cache_key(etcd_key, conn._url)
        if cache in self.__local_cache:
            for field in self.__local_cache[cache]:
                field.set_value(key, value)
        return conn.put(etcd_key, repr(value))


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
            conn = next(config.conn() for config in configs)
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
        self.env_name = env_name or self.env_name

        def _config(conf: type):
            ignore = ("__dict__", "__weakref__")
            # noinspection PyTypeChecker
            new_conf: Type[EtcdConfig] = type(
                conf.__name__,
                (*conf.mro()[:-1], EtcdConfig, conf.mro()[-1]),
                {key: val for key, val in conf.__dict__.items() if key not in ignore},
            )

            @wraps(new_conf)
            def _wrap(*args, **kwargs) -> EtcdConfig:
                new_conf.env_name = env_name or self.env_name
                new_conf.conn = proxy(self.conn)
                new_conf.data_access = self.data_access
                ins = new_conf(*args, **kwargs)  # type: ignore
                self.data_access.register_listener(ins)
                return ins

            return _wrap

        return _config

    def put(self, key: str, value: Any):
        self.data_access.put(key, value, self.env_name, self.conn)

    def get(self, key: str) -> ConfigField:
        return self.data_access.get(key, env_name=self.env_name, conn=self.conn)

from ctypes import cdll, c_char_p, c_long, c_int
import os
from microkit import PACKAGE_PATH


class Consul:
    _consul_so = cdll.LoadLibrary(os.path.join(PACKAGE_PATH, "libs/consul.so"))

    # method: consulFactory
    _consul_so.consulFactory.argtype = c_char_p  # type: ignore
    _consul_so.consulFactory.restype = c_long

    # method: consulRegister
    _consul_so.consulRegister.argtypes = [
        c_long,
        c_char_p,
        c_char_p,
        c_char_p,
        c_int,
        c_char_p,
    ]

    # method: consulDeregister
    _consul_so.consulDeregister.argtypes = [c_long, c_char_p]

    def __init__(self, address: str):
        self._consul: c_long = self._consul_so.consulFactory(address.encode("UTF8"))

    def register(
        self,
        service_id: str,
        service_name: str,
        service_address: str,
        server_port: int,
        tags: str = "",
    ):
        """
        向consul中注册服务
        Args:
            service_id: 服务id，标识
            service_name: 服务名字
            ...
        """
        self._consul_so.consulRegister(
            self._consul,
            service_id.encode("UTF8"),
            service_name.encode("UTF8"),
            service_address.encode("UTF8"),
            server_port,
            tags.encode("UTF8"),
        )

    def deregister(self, service_id: str):
        """
        为服务进行反注册
        Args:
            service_id: 服务id
        """
        self._consul_so.consulDeregister(self._consul, service_id.encode("UTF8"))
from ctypes import c_long, c_char_p, c_int

def consulFactory(address: str) -> c_long: ...
def consulRegister(
    consul: c_long,
    service_id: c_char_p,
    service_name: c_char_p,
    service_address: c_char_p,
    server_port: c_int,
    tags: c_char_p,
): ...
def consulDeregister(consul: c_long, service_id: c_char_p): ...
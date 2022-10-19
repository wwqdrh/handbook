import struct
import binascii
import ctypes

values = (1, b"good", 1.22)  # 查看格式化字符串可知，字符串必须为字节流类型。
s = struct.struct("I4sf")
buff = ctypes.create_string_buffer(s.size)
packed_data = s.pack_into(buff, 0, *values)
unpacked_data = s.unpack_from(buff, 0)

print("Original values:", values)
print("Format string :", s.format)
print("buff :", buff)
print("Packed Value :", binascii.hexlify(buff))
print("Unpacked Type :", type(unpacked_data), " Value:", unpacked_data)

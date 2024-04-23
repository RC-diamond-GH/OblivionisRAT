def bytes_to_int(byte_string):
    # 将字节串转换为整数
    num = int.from_bytes(byte_string, byteorder='big')
    return num

byte_string= b"\xf0\x0d\xbe\xef\xde\xad\xbe\xef\xf0\x0d\xbe\xef\xde\xad\xbe\xef"
result = bytes_to_int(byte_string)
print("Result:", result)

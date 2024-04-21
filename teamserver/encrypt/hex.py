def hex_to_string(hex_value):
    hex_string = hex_value[2:]  # 去掉开头的'0x'
    result = '\\x' + '\\x'.join(hex_string[i:i+2] for i in range(0, len(hex_string), 2))
    return result

hex_value = 0xf00dbeefdeadbeeff00dbeefdeadbeef
formatted_string = hex_to_string(hex(hex_value))
print(formatted_string)


def string_to_hex(input_string):
    hex_string = input_string.replace('\\x', '')
    return int(hex_string, 16)

def string_to_hex_format(input_string, format):
    hex_string = input_string.replace('\\x', format)
    return hex_string

formatted_string = r'\xf0\x0d\xbe\xef\xde\xad\xbe\xef\xf0\x0d\xbe\xef\xde\xad\xbe\xef'

hex_value = string_to_hex(formatted_string)
print(hex(hex_value))

formatted_hex_string = string_to_hex_format(formatted_string, ',0x')
print(formatted_hex_string)

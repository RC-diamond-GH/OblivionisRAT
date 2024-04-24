import requests
import random
import json

def rc4(key, text):
    S = list(range(256))
    j = 0

    # KSA (Key Scheduling Algorithm)
    for i in range(256):
        j = (j + S[i] + key[i % len(key)]) % 256
        S[i], S[j] = S[j], S[i]

    # PRGA (Pseudo-Random Generation Algorithm)
    i = j = 0
    result = []
    for char in text:
        i = (i + 1) % 256
        j = (j + S[i]) % 256
        S[i], S[j] = S[j], S[i]
        result.append(char ^ S[(S[i] + S[j]) % 256])
    return bytes(result)




def string_to_byte_stream(data):
    byte_stream = data.encode()
    padded_stream = byte_stream.ljust(16, b'\x00')
    return padded_stream


def print_byte_stream(byte_stream):
    hex_string = ' '.join(format(byte, '02x') for byte in byte_stream)
    print(hex_string)

def int_to_bytes(num):
    # 计算所需字节数
    num_bytes = (num.bit_length() + 7) // 8
    # 转换为字节流
    return num.to_bytes(num_bytes, 'big')

def mod_pow(u, amount):
    u_bytes = string_to_byte_stream(u)

    u_int = int.from_bytes(u_bytes, byteorder='big')
    mod = 2**127 - 1
    result = pow(u_int, amount, mod)
    return result

def mod_pow2(u, amount):
    mod = 2**127 - 1
    result = pow(u, amount, mod)
    return result

def send_get(url, headers=None, cookies=None):
    try:
        response = requests.get(url, headers=headers, cookies=cookies)
    except requests.exceptions.RequestException as e:
        print("GET Request Error:", e)


def send_post(url, headers, byte_stream):
    try:
        response = requests.post(url, headers=headers, data=byte_stream)
        return bytes(reversed(response.content[-16:]))
    except Exception as e:
        print("Error:", e)


def json_to_bytes(json_data):
    try:
        # 将 JSON 数据转换为字节流
        json_bytes = json.dumps(json_data).encode('utf-8')
        return json_bytes
    except Exception as e:
        print("Error converting JSON to bytes:", str(e))
        return None

url = "http://localhost:8080"
headers = {'Custom-Header1': 'Value1', 'Custom-Header2': 'Value2'}
cookies = {'cookie1': 'FUIo3qOUftdVX9XOHFVXCxBMPkEeQEz3cwVWr+VPTJQ='}
byte_stream = b"\xf0\x0d\xbe\xef\xde\xad\xbe\xef\xf0\x0d\xbe\xef\xde\xad\xbe\xef"

send_get(url,headers, cookies)

u = "USAnmslhahahahah"
amount = random.randint(0, 2**128 - 1)  # 生成一个128位的随机数
result1 = mod_pow(u, amount)

srvAes = send_post(url,headers,result1.to_bytes((result1.bit_length() + 7) // 8, byteorder='big'))


aeskey = mod_pow2(int.from_bytes(srvAes, byteorder='big'), amount)

print(aeskey)


json_data = {
    "arch": "x86_64",
    "hostname": "example.com",
    "system": "Linux"
}

keyjson = rc4(int_to_bytes(aeskey), json_to_bytes(json_data))

send_post(url, headers, keyjson)








json_data = {
    "arch": "x86_64",
    "hostname": "example.com",
    "system": "Linux"
}

json_bytes = json.dumps(json_data).encode('utf-8')



#319086091509812015005475032388852367087
#319086091509812015005475032388852367087







'''
def send_get_with_bytes(url, headers, byte_stream):
    try:
        response = requests.get(url, headers=headers)
        print("Response:", response.text)
    except Exception as e:
        print("Error:", e)

url = "http://127.0.0.1:8080"
headers = {"Custom-Header1": "Value1", "Custom-Header2": "Value2"}
cookies = {'cookie1': 'value1'}
byte_stream = b"\xf0\x0d\xbe\xef\xde\xad\xbe\xef\xf0\x0d\xbe\xef\xde\xad\xbe\xef"

send_post_with_bytes(url, headers, byte_stream)
'''
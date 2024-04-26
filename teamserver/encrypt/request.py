import requests
import time

def send_post_request(url, headers, data):
    try:
        response = requests.post(url, headers=headers, data=data)
        print("Response:", response.text)
    except Exception as e:
        print("Error:", e)

if __name__ == "__main__":
    # 设置目标IP和端口
    ip_address = "127.0.0.1"
    port = "50050"

    # 设置请求头和请求体
    headers = {
        "Username": "admin",
        "Password": "password",
        "User-Agent": "Value1",
    }

    # 构造请求URL
    url = "http://" + ip_address + ":" + port + "/client"

    # 每秒发送请求
    while True:
        # 用户输入字节流数据
        byte_data = input("请输入字节流数据（以逗号分隔）：")

        # 将用户输入的数据转换为字节流
        try:
            data = bytes.fromhex(byte_data.replace(" ", "").replace(",", ""))
        except ValueError:
            print("输入的数据不是有效的字节流，请重新输入。")
            continue

        # 发送请求
        send_post_request(url, headers, data)
        time.sleep(1)

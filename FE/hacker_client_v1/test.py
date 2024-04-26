import socket

def main():
    # 设置服务器地址和端口
    host = "127.0.0.1"
    port = 8080

    # 创建 TCP 套接字
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
        # 连接到服务器
        client_socket.connect((host, port))

        # 发送数据给服务器
        client_socket.sendall(b"hello server")

        # 接收服务器返回的数据
        data = client_socket.recv(1024)

        # 打印服务器返回的数据
        print(f"Received: {data.decode()}")

if __name__ == "__main__":
    main()

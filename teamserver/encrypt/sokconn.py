import websocket
import threading
import time

def on_message(ws, message):
    print(f"Received message: {message}")

def on_error(ws, error):
    print(f"Error: {error}")

def on_close(ws):
    print("WebSocket connection closed")

def on_open(ws):
    print("WebSocket connection established")

    # 每隔一秒发送一次消息
    def send_message():
        while True:
            time.sleep(1)
            ws.send("Hello, server!")

    # 启动发送消息的线程
    threading.Thread(target=send_message).start()

if __name__ == "__main__":
    # WebSocket 服务器地址
    websocket_url = "ws://localhost:8080"

    # 设置请求头
    header = {
        "Custom-Header1": "Value1",
        "Custom-Header2": "Value2",
    }

    # 设置 Cookie
    cookie = "session_id=123456789"

    # 创建 WebSocket 连接并传递请求头和 Cookie
    ws = websocket.create_connection(websocket_url,
                                     header=header,
                                     cookie=cookie,
                                     on_message=on_message,
                                     on_error=on_error,
                                     on_close=on_close,
                                     on_open=on_open)

    # 运行 WebSocket 连接
    ws.run_forever()

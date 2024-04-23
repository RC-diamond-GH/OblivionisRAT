package main

import (
	"encoding/base64"
	"fmt"
	"reflect"
)

func GetBytes(data []byte, length int) []byte {
	if length > len(data) {
		length = len(data)
	}
	data = data[:length]
	return data
}

func GET_handler(cookie string, listener Listener) ([]byte, bool) {
	var res []byte
	cookie_decode, err := base64.StdEncoding.DecodeString(cookie)
	aAES := getAES(listener.A)

	if err != nil {
		fmt.Println("AESa base64 not match" + err.Error())
		return res, false
	} else if len(cookie_decode) != 32 {
		fmt.Println("AESa 16BYTES not match")

		return res, false
	} else if reflect.DeepEqual(listener.A, aAES.DecryptData(cookie_decode)) {
		fmt.Println("AESa had match")
		return res, true
	} else {
		fmt.Println("AESa not match")
		return res, false
	}
}

func POST_handler(body []byte, listener Listener) []byte {
	var res []byte
	if len(body) == 16 {
		// regis beacon
		xmlname := bToHexString(body)
		if checkXMLExists("./data/", xmlname+".xml") {
			fmt.Println("have the xml")
		} else {
			beacons := []Beacon{
				{Hostname: "example.com", Ip: "192.168.1.1", Domin: "example", CusAES: "abcd", AESkey: "123456", Live: true},
				{Hostname: "test.com", Ip: "192.168.1.2", Domin: "test", CusAES: "efgh", AESkey: "789012", Live: false},
			}
			listener := Listener{
				Lisname: "listener1",
				Port:    8080,
				A:       convertToUint8("usanmsl"),
				Beacons: beacons,
			}

			err := saveXML(xmlname+".xml", listener)
			if err != nil {
				fmt.Println("can not save xml")
			}
		}
	}
	return res
}

func printkey(key []uint8) {
	for _, b := range key {
		fmt.Printf("%x ", b)
	}
}

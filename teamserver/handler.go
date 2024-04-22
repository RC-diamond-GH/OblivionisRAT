package main

import (
	"encoding/base64"
	"fmt"
)

func GetBytes(data []byte, length int) []byte {
	if length > len(data) {
		length = len(data)
	}
	data = data[:length]
	return data
}

func GET_handler(cookie string) []byte {
	var res []byte
	temp_key := []uint8{0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12, 0x12}
	aAES := getAES(temp_key)
	//	tempe := base64.StdEncoding.EncodeToString(aAES.EncryptData(temp_key))
	// println(tempe)

	undecodeCusAES, _ := base64.StdEncoding.DecodeString(cookie)
	printkey(aAES.DecryptData(undecodeCusAES))
	//aAES := getAES(miyao)
	//aAES.DecryptData(jiemishuju)
	return res
}

func POST_handler(body []byte) []byte {
	var res []byte
	if len(body) == 16 {
		// regis beacon
		xmlname := bToHexString(body)
		if checkXMLExists(xmlname + ".xml") {
			fmt.Println("have the xml")
		} else {
			beacons := []Beacon{
				{Hostname: "example.com", Ip: "192.168.1.1", Domin: "example", CusAES: "abcd", AESkey: "123456", Live: true},
				{Hostname: "test.com", Ip: "192.168.1.2", Domin: "test", CusAES: "efgh", AESkey: "789012", Live: false},
			}
			listener := Listener{
				Lisname: "listener1",
				Port:    8080,
				A:       123,
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

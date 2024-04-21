package main

import (
	"fmt"
)

func GetBytes(data []byte, length int) []byte {
	if length > len(data) {
		length = len(data)
	}
	data = data[:length]
	return data
}

func POST_handler(body []byte) []byte {
	var res []byte
	if len(body) == 16 {
		// regis beacon
		xmlname := bToHexString(body)
		if checkXMLExists(xmlname + ".xml") {
			fmt.Println("have the xml")
		} else {
			beacon := Beacon{
				hostname: "123",
				ip:       "123",
				domin:    "123",
				aeskey:   "123",
				live:     true,
			}
			err := saveXML(xmlname+".xml", beacon)
			if err != nil {
				fmt.Println("can not save xml")
			}
		}
	}
	return res
}

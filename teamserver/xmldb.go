package main

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Beacon struct {
	Hostname string `xml:"hostname"`
	Ip       string `xml:"ip"`
	Domin    string `xml:"domin"`
	CusAES   string `xml:"CusAES"`
	AESkey   string `xml:"AESkey"`
	Live     bool   `xml:"live"`
}
type Listener struct {
	Lisname string   `xml:"lisname"`
	Port    int      `xml:"port"`
	A       int      `xml:"a"`
	Beacons []Beacon `xml:"beacon"`
}

func saveXML(filename string, data Listener) error {
	filename = "./data/" + filename
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "    ")
	if err := encoder.Encode(data); err != nil {
		return err
	}
	if err := encoder.Flush(); err != nil {
		panic(err)
	}
	return nil
}

func readXML(filename string, data Beacon) error {
	filename = "./data/" + filename
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		return err
	}
	return nil
}

func checkXMLExists(filename string) bool {
	dir, _ := ioutil.ReadDir("./data")

	for _, file := range dir {
		if file.Name() == filename {
			return true
		}
	}

	return false
}

func bToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func bytesToHexString(data []byte) string {
	hexString := ""
	for _, b := range data {
		hexString += fmt.Sprintf("\\x%02x", b)
	}
	return hexString
}

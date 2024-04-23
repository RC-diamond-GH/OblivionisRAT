package main

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

type Beacon struct {
	Hostname string `xml:"hostname"`
	Ip       string `xml:"ip"`
	Domin    string `xml:"domin"`
	Arch     string `xml:"arch"`
	System   string `xml:"system"`
	CusAES   string `xml:"CusAES"`
	AESkey   string `xml:"AESkey"`
	Live     bool   `xml:"live"`
}
type Listener struct {
	Lisname string   `xml:"lisname"`
	Uri     string   `xml:"uri"`
	Port    int      `xml:"port"`
	A       []uint8  `xml:"a"`
	Beacons []Beacon `xml:"beacon"`
}

func saveXML(filename string, data *Listener) error {
	filename = filename + ".xml"
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

func checkXMLExists(checkpath string, filename string) bool {
	dir, _ := ioutil.ReadDir(checkpath)

	for _, file := range dir {
		if file.Name() == filename {
			return true
		}
	}

	return false
}

func ReadXML(filename string, obj interface{}) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	return nil
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

func Check_Beacon_CusAES(listener *Listener, cusaes big.Int) bool {
	for _, beacon := range listener.Beacons {
		if beacon.CusAES == cusaes.String() {
			return true
		}
	}
	return false
}

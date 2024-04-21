package main

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Beacon struct {
	hostname string `xml:"hostname"`
	ip       string `xml:"ip"`
	domin    string `xml:"domin"`
	aeskey   string `xml:"aeskey"`
	live     bool   `xml:"live"`
}

func saveXML(filename string, data Beacon) error {
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

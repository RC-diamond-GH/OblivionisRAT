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

func Listener_Handler(w http.ResponseWriter, r *http.Request, listener *Listener) {
	headers := r.Header
	var res []byte

	for header, expectedValue := range expectedHeaders {
		value := headers.Get(header)

		if value != expectedValue {
			http.Error(w, fmt.Sprintf("Forbidden: Header value for %s does not match", header), http.StatusForbidden)
			return
		}
	}

	if r.Method == http.MethodGet {
		cookies := r.Cookies()
		cookie := ""

		for _, i := range cookies {
			cookie += i.Value
		}
		res, love := GET_handler(cookie, listener)
		if !love {
			http.Error(w, fmt.Sprintf("Forbidden: Cookie for %s does not match", cookie), http.StatusForbidden)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			res_tmp := make([]byte, 4)
			binary.LittleEndian.PutUint32(res_tmp, 0xbeebeebe)
			res = append(res, res_tmp...)
			w.Write(res)
			res = make([]byte, 0)
		}

	}

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		w.WriteHeader(http.StatusOK)

		res = POST_handler(body, listener, r)

		res_tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(res_tmp, 0xbeebeebe)
		res = append(res, res_tmp...)

		w.Write(res)
		res = make([]byte, 0)
	}
}

func main() {
	uri := ""
	port := 8080
	lisname := "ilovec2"
	StartListener(uri, port, lisname)
}

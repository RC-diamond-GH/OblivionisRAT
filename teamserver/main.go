package main

import (
	"fmt"
	"net/http"
)

var expectedHeaders = map[string]string{
	"Custom-Header1": "Value1",
	"Custom-Header2": "Value2",
}

func handler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header

	for header, expectedValue := range expectedHeaders {
		value := headers.Get(header)

		if value != expectedValue {
			http.Error(w, fmt.Sprintf("Forbidden: Header value for %s does not match", header), http.StatusForbidden)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Request accepted")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

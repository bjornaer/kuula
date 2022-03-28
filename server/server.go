package main

import (
	"fmt"
	"io"
	"net/http"
)

func dcd(w http.ResponseWriter, req *http.Request) {
	bod, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "error\n")
	}
	fmt.Println(string(bod))

	fmt.Fprintf(w, "done.\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/decode", dcd)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

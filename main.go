package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/shorted", Add_data)
	addr := ":80"
	fmt.Println("Server served on ", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

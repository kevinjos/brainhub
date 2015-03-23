package main

import (
	"flag"
	"fmt"
	"net/http"
)

var addr = flag.String("addr", ":8889", "http service address")

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println("Serving from addr" + *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic("ListenAndServe:" + err.Error())
	}
}

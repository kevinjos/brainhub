package main

import (
	"net/http"
	"html/template"
	"flag"
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
	rootTempl := template.Must(template.ParseFiles("static/index.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	rootTempl.Execute(w, r.Host)
}

func main() {

	shutdown := make(chan bool)
	http.HandleFunc("/", rootHandler)
	go http.ListenAndServe(*addr, nil)
	<-shutdown

}

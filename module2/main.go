package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	for k, v := range header {
		w.Header().Set(k, v[0])
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	log.Println("client ip:", r.RemoteAddr)
	io.WriteString(w, "ok")
}

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

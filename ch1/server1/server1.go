package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got request from %s", r.RemoteAddr)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
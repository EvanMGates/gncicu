package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	createServer()
}

func createServer() {
	hs := &http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/", rootHandler)
	log.Fatal(hs.ListenAndServe())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "No. I see you, Game Night Crew.")
}

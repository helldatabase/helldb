package server

import (
	"log"
	"net/http"
)

func Serve(port string) {

	http.HandleFunc("/query", query)
	http.HandleFunc("/status", status)
	http.HandleFunc("/length", length)

	log.Printf("serving hell on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

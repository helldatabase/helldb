package server

import (
	"log"
	"net/http"
)

func Serve(port string) {

	http.HandleFunc("/query", query)

	log.Printf("serving hell on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

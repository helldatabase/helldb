package server

import (
	"fmt"
	"log"
	"net/http"
)

func Serve(port string) {

	http.HandleFunc("/query", query)

	fmt.Printf("serving hell on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

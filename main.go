package main

import (
	"flag"
	"os"
	"strconv"

	"helldb/portal/evaluator"
	"helldb/portal/server"
)

func main() {

	clientFlag := flag.NewFlagSet("client", flag.ExitOnError)
	serverFlag := flag.NewFlagSet("server", flag.ExitOnError)

	serverPortPtr := serverFlag.Uint("port", 8080, "port to run helldb on")
	clientPromptPtr := clientFlag.String("prompt", ">>> ", "prompt to use for client")

	flag.Parse()

	if len(os.Args) <= 1 {
		server.Serve(strconv.Itoa(int(*serverPortPtr)))
	}

	switch os.Args[1] {
	case "server":
		server.Serve(strconv.Itoa(int(*serverPortPtr)))
	case "client":
		evaluator.REPL(*clientPromptPtr)
	}

}

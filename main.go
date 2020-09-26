package main

import (
	"flag"
	"os"
	"strconv"

	"helldb/portal/evaluator"
	"helldb/portal/server"
)

func main() {

	const (
		version   = "1.0.0"
		debugStr  = "debug"
		serverStr = "server"
	)

	debugFlag := flag.NewFlagSet(debugStr, flag.ExitOnError)
	serverFlag := flag.NewFlagSet(serverStr, flag.ExitOnError)
	versionPtr := flag.Bool("v", false, "displays current database version")

	serverPortPtr := serverFlag.Uint("port", 8080, "port to run helldb on")
	debugPromptPtr := debugFlag.String("prompt", ">>> ", "prompt to use for client")

	flag.Parse()

	if *versionPtr {
		println(version)
		os.Exit(0)
	}

	if len(os.Args) <= 1 {
		server.Serve(strconv.Itoa(int(*serverPortPtr)))
	}

	switch os.Args[1] {
	case debugStr:
		evaluator.REPL(*debugPromptPtr)
	case serverStr:
		server.Serve(strconv.Itoa(int(*serverPortPtr)))
	}

}

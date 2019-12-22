package main

import (
	"flag"
	"fmt"
	"net/dial/ioevents"
	Port "net/dial/port"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: test-local-port [port number]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

const BEGIN_PORT = 5000
const END_PORT = 8080

func main() {
	flag.Usage = usage
	flag.Parse()

	pathFile := "ports.json"
	key := "experiment1"

	if port, err := ioevents.Read(pathFile, key); err == nil {
		fmt.Println(port)
		return
	}

	port, err := Port.FindFreePort(BEGIN_PORT, END_PORT)
	ioevents.Write(pathFile, key, port)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(port)
}

package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: test-local-port [port number]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func checkPort(port string) (string, error) {
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid port %q: %s", port, err)
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		errorText := fmt.Sprintf("Can't listen on port %q: %s", port, err)
		return "", errors.New(errorText)
	}

	err = ln.Close()
	if err != nil {
		errorText := fmt.Sprintf("Couldn't stop listening on port %q: %s", port, err)
		return "", errors.New(errorText)
	}

	return port, nil
}

func findFreePort(beginPort int, endPort int) (int, error) {
	for port := beginPort; port <= endPort; port++ {
		strPort := strconv.FormatInt(int64(port), 10)
		_, err := checkPort(strPort)

		if err == nil {
			fmt.Printf("TCP Port %q is available", strPort)
			return port, nil
		}
	}
	return 0, errors.New("not free ports")
}

func main() {
	const BEGIN_PORT = 5000
	const END_PORT = 8080
	flag.Usage = usage
	flag.Parse()

	port, err := findFreePort(BEGIN_PORT, END_PORT)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(port)
}

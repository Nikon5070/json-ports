package main

import (
	"flag"
	"fmt"
	Port "github.com/Nikon5070/json-ports/port"
	"os"
)

const BEGIN_PORT = 5000
const END_PORT = 8080

func main() {
	var jp Port.JsonPort
	var pathFile, name string

	flag.StringVar(&pathFile, "file", "./ports.json", "file with busy name branches and ports")
	flag.StringVar(&name, "name", "master", "name branch")

	flag.Parse()

	if port, err := Port.Read(pathFile, name, &jp); err == nil {
		fmt.Println(port)
		return
	}

	port, err := Port.FindFreePort(BEGIN_PORT, END_PORT, jp)

	Port.Write(pathFile, name, port, jp)
	fmt.Println(jp)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(port)
}

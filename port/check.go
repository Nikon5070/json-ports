package port

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
)

func CheckPort(port string) (string, error) {
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

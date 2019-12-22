package port

import (
	"errors"
	"strconv"
)

func FindFreePort(beginPort int, endPort int) (int, error) {
	for port := beginPort; port <= endPort; port++ {
		strPort := strconv.FormatInt(int64(port), 10)
		_, err := CheckPort(strPort)

		if err == nil {
			return port, nil
		}
	}
	return 0, errors.New("not free ports")
}

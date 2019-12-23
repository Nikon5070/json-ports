package port

import (
	"errors"
	"strconv"
)

func GetBusyPortFromFile(jm JsonPort) []int {
	var array []int
	for _, v := range jm {
		array = append(array, v)
	}
	return array
}

func FindElementInArray(a []int, i int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func FindFreePort(beginPort int, endPort int, jp JsonPort) (int, error) {
	busyPorts := GetBusyPortFromFile(jp)
	for port := beginPort; port <= endPort; port++ {

		if FindElementInArray(busyPorts, port) {
			continue
		}

		strPort := strconv.FormatInt(int64(port), 10)
		_, err := CheckPort(strPort)

		if err == nil {
			return port, nil
		}
	}
	return 0, errors.New("not free ports")
}

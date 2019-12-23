package port

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadFile(pathFile string) string {
	file, err := ioutil.ReadFile(pathFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(file)
}

func ConvertFileJsonToMap(pathFile string) (map[string]interface{}, error) {
	jsonStr := ReadFile(pathFile)
	jsonMap := make(map[string]interface{})

	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		return nil, err
	}
	return jsonMap, nil
}

func convertJson(m map[string]interface{}) JsonPort {
	jp := make(JsonPort)
	for k, v := range m {
		jp[k] = int(v.(float64))
	}
	return jp
}

func Read(pathFile, key string, jp *JsonPort) (int, error) {
	m, err := ConvertFileJsonToMap(pathFile)
	if err != nil {
		return 0, err
	}

	*jp = convertJson(m)

	value := m[key]
	if value == nil {
		err := errors.New("not found this key")
		return 0, err
	}

	strValue := fmt.Sprintf("%v", value)
	if v, err := strconv.Atoi(strValue); err == nil {
		return v, nil
	} else {
		return 0, err
	}
}

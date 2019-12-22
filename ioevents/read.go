package ioevents

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

func Read(pathFile, key string) (int, error) {

	jsonMap, err := ConvertFileJsonToMap(pathFile)
	if err != nil {
		return 0, err
	}

	value := jsonMap[key]
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

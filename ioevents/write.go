package ioevents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Write(pathFile, key string, value int) {
	jsonMap, err := ConvertFileJsonToMap(pathFile)
	if err != nil {
		return
	}

	jsonMap[key] = value
	jsonMarshal, _ := json.MarshalIndent(jsonMap, "", "    ")
	fmt.Println(string(jsonMarshal))

	if err = ioutil.WriteFile(pathFile, jsonMarshal, 0777); err != nil {
		return
	}
}

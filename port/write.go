package port

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Write(pathFile, key string, value int, jp JsonPort) {
	jp[key] = value
	jsonMarshal, _ := json.MarshalIndent(jp, "", "    ")
	fmt.Println(string(jsonMarshal))

	if err := ioutil.WriteFile(pathFile, jsonMarshal, 0755); err != nil {
		return
	}
}

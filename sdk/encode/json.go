package main

import (
	"encoding/json"
	"fmt"
)

type rawData struct {
	Data string `json:"data"`
}

//支持自定义对象在进行json.Marshal时进行重载
func (r *rawData) MarshalJSON() ([]byte, error) {
	return []byte(nil), nil
}

func StudyJson() {
	raw1Data := rawData{"data"}
	json1Data, _ := json.Marshal(raw1Data)
	fmt.Println(string(json1Data), raw1Data)
	raw2Data := rawData{}
	err := json.Unmarshal(json1Data, &raw2Data)
	if err == nil {
		fmt.Println(raw2Data)
	}

	raw3Data := map[string]int{
		"data": 1,
	}

	json3Data, _ := json.Marshal(raw3Data)

	fmt.Println(raw3Data)
	fmt.Println(string(json3Data))
}

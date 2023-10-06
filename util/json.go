package util

import (
	"encoding/json"
	"fmt"
)

func JsonMarshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("`JsonMarshal failed for v: %+v", v)
	}
	return string(b)
}

func JsonUnmarshal(jsonStr string, v interface{}) error {
	byteValue := []byte(jsonStr)
	err := json.Unmarshal(byteValue, v)
	if err != nil {
		return fmt.Errorf("`JsonUmarshal failed for v: %+v", v)
	}
	return nil
}

func ConvertEntity(source interface{}, target interface{}) error {
	jsonStr := JsonMarshal(source)
	return JsonUnmarshal(jsonStr, target)
}

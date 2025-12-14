package appCommon

import (
	"github.com/goccy/go-json"
)

func MarshalData(value interface{}) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func JsonToString(value ...interface{}) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func StringToJson(value string, target interface{}) error {
	return json.Unmarshal([]byte(value), target)
}

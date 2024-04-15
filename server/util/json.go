package util

import json "github.com/json-iterator/go"

func UnmarshalConverter(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}

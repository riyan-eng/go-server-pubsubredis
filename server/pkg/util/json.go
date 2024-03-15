package util

import json "github.com/json-iterator/go"

func UnmarshalConverter[T any](s string) (data T) {
	json.Unmarshal([]byte(s), &data)
	return
}

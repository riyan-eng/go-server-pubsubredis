package util

import "encoding/json"

func UnmarshalConverter[T any](s []byte) (data T) {
	json.Unmarshal(s, &data)
	return
}

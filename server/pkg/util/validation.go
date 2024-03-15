package util

func IsValid(input interface{}) bool {
	switch v := input.(type) {
	case int:
		return v != 0
	case string:
		return v != ""
	default:
		return false
	}
}

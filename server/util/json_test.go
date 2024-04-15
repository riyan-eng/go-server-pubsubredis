package util

import (
	"testing"
)

func TestUnmarshalConverter(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Create a JSON string
	jsonStr := `{"name": "John", "age": 30}`

	// Unmarshal the JSON string into a TestStruct
	var testStruct TestStruct
	err := UnmarshalConverter(jsonStr, &testStruct)
	if err != nil {
		t.Errorf("UnmarshalConverter() failed: %v", err)
	}

	// Check the values of the TestStruct
	if testStruct.Name != "John" {
		t.Errorf("UnmarshalConverter() failed: expected Name to be 'John', got '%s'", testStruct.Name)
	}
	if testStruct.Age != 30 {
		t.Errorf("UnmarshalConverter() failed: expected Age to be 30, got %d", testStruct.Age)
	}
}

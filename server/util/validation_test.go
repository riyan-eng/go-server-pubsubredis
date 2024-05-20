package util

import (
	"testing"
)

func TestIsValidStruct_String(t *testing.T) {
	validator := NewIsValid()
	if !validator.String("hello") {
		t.Error("Expected true for non-empty string")
	}
	if validator.String("") {
		t.Error("Expected false for empty string")
	}
}

func TestIsValidStruct_Int(t *testing.T) {
	validator := NewIsValid()
	if !validator.Int(1) {
		t.Error("Expected true for non-zero int")
	}
	if validator.Int(0) {
		t.Error("Expected false for zero int")
	}
}

func TestIsValidStruct_Float64(t *testing.T) {
	validator := NewIsValid()
	if !validator.Float64(1.1) {
		t.Error("Expected true for non-zero float64")
	}
	if validator.Float64(0.0) {
		t.Error("Expected false for zero float64")
	}
}

func TestIsValidStruct_Float32(t *testing.T) {
	validator := NewIsValid()
	if !validator.Float32(1.1) {
		t.Error("Expected true for non-zero float32")
	}
	if validator.Float32(0.0) {
		t.Error("Expected false for zero float32")
	}
}

func TestIsValidStruct_Int64(t *testing.T) {
	validator := NewIsValid()
	if !validator.Int64(1) {
		t.Error("Expected true for non-zero int64")
	}
	if validator.Int64(0) {
		t.Error("Expected false for zero int64")
	}
}

func TestIsValidStruct_Int32(t *testing.T) {
	validator := NewIsValid()
	if !validator.Int32(1) {
		t.Error("Expected true for non-zero int32")
	}
	if validator.Int32(0) {
		t.Error("Expected false for zero int32")
	}
}

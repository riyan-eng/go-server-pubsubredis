package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyToString(t *testing.T) {
	c := Convert()
	// Test with a string
	var str any = "test"
	assert.Equal(t, "test", c.AnyToString(str))
}

func TestAnyToInt(t *testing.T) {
	c := Convert()
	// Test with an int
	var num any = 123
	assert.Equal(t, 123, c.AnyToInt(num))
}

func TestAnyToFloat32(t *testing.T) {
	c := Convert()
	// Test with a float
	var flt any = float32(123.456)
	assert.Equal(t, float32(123.456), c.AnyToFloat32(flt))
}

func TestAnyToFloat64(t *testing.T) {
	c := Convert()
	// Test with a float
	var flt any = 123.456
	assert.Equal(t, 123.456, c.AnyToFloat64(flt))
}

func TestAnyToBool(t *testing.T) {
	c := Convert()
	// Test with a bool
	var bl any = true
	assert.Equal(t, true, c.AnyToBool(bl))
}

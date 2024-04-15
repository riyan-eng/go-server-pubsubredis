package util

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	str := "password"
	hashedStr := GenerateHash(str)

	if hashedStr == str {
		t.Errorf("GenerateHash(%s) = %s; want hashed string", str, hashedStr)
	}
}

func TestVerifyHash(t *testing.T) {
	str := "password"
	hashedStr := GenerateHash(str)

	if !VerifyHash(hashedStr, str) {
		t.Errorf("VerifyHash(%s, %s) = false; want true", hashedStr, str)
	}
}

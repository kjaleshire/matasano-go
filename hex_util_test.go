package matasano_test

import (
	matasano "github.com/kjaleshire/matasano-go"
	"testing"
)

func TestHexCharsToValuesBase64(t *testing.T) {
	b64, err := matasano.HexCharsToValuesBase64(HexString)
	if err != nil || b64 != Base64String {
		t.Error("Failed to convert string")
	}
}

func TestStringXor(t *testing.T) {
	hex_string, err := matasano.StringXor(XorString1, XorString2)
	if err != nil {
		t.Error("Failed to convert string")
	} else if hex_string != XorAnswer {
		t.Errorf("Wrong answer, got %s", hex_string)
	}
}

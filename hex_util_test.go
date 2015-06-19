package matasano_test

import (
	matasano "github.com/kjaleshire/matasano-go"
	"testing"
)

// Challenge 1 test
func TestHexCharsToValuesBase64(t *testing.T) {
	b64 := matasano.HexCharsToValuesBase64(HexString)
	if b64 != Base64String {
		t.Errorf("Wrong answer, got %s", b64)
	}
}

// Challenge 2 test
func TestStringXor(t *testing.T) {
	hex_string, err := matasano.StringXor(XorString1, XorString2)
	if err != nil {
		t.Error("Failed to convert string")
	} else if hex_string != XorAnswer {
		t.Errorf("Wrong answer, got %s", hex_string)
	}
}

// Challenge 3 test
func TestSingleByteKeyDecoding(t *testing.T) {
	decode := matasano.BreakSingleByteCipher(EncodedString)
	if decode.String != DecodedString {
		t.Errorf("Wrong answer, got %s", decode.String)
	}
}

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
	hexString := matasano.StringXor(XorString1, XorString2)
	if hexString != XorAnswer {
		t.Errorf("Wrong answer, got %s", hexString)
	}
}

// Challenge 3 test
func TestSingleByteSingleStringDetecting(t *testing.T) {
	decode := matasano.BreakSingleLineByteCipher(EncodedString)
	if decode.String != DecodedString || decode.Cipher != DecodedCipher {
		t.Errorf("Wrong answer, got %s with cipher 0x%x", decode.String, decode.Cipher)
	}
}

// Challenge 3 test
func TestSingleByteMultiStringDetecting(t *testing.T) {
	decode := matasano.BreakMultiLineFileByteCipher(FilePath)
	if decode.String != DecodedFileString || decode.Cipher != DecodedFileCipher || decode.Line != DecodedFileLine {
		t.Errorf("Wrong answer, got %s on line %d with cipher 0x%x", decode.String, decode.Line, decode.Cipher)
	}
}

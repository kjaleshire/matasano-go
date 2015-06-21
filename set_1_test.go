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
func TestSingleStringByteCipherDetecting(t *testing.T) {
	valueBytes := matasano.HexCharsToValues(EncodedString)
	decode := matasano.BreakSingleLineByteCipher(valueBytes)
	if decode.String != DecodedString || decode.Cipher != DecodedCipher {
		t.Errorf("Wrong answer, got %s with cipher 0x%x", decode.String, decode.Cipher)
	}
}

// Challenge 4 test
func TestMultiStringFileByteCipherDetecting(t *testing.T) {
	decode := matasano.BreakMultiLineFileByteCipher(FilePath)
	if decode.String != DecodedFileString || decode.Cipher != DecodedFileCipher || decode.Line != DecodedFileLine {
		t.Errorf("Wrong answer, got %s on line %d with cipher 0x%x", decode.String, decode.Line, decode.Cipher)
	}
}

// Challenge 5 test
func TestReapeatingKeyXor(t *testing.T) {
	encodedString := matasano.RepeatingKeyXor([]byte(OpeningStanza), []byte(RepeatingKeyCipher))
	if encodedString != RepeatingXorResult {
		t.Errorf("Wrong answer, got %s", encodedString)
	}
}

// Challenge 6 tests
func TestHammingDistance(t *testing.T) {
	distance := matasano.HammingBitDistance([]byte(HammingString1), []byte(HammingString2))
	if distance != HammingDistance {
		t.Errorf("Wrong answer, got %d", distance)
	}
}

func TestBreakRepeatingKeyXorFile(t *testing.T) {
	decode := matasano.BreakRepeatingKeyXorFile(RepeatingKeyXorFilePath)
	if decode.Key != RepeatingXorKey {
		t.Errorf("Wrong answer, got key %s\nSize: %d", decode.Key, decode.KeySize)
	}
}

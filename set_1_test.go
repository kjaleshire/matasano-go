package matasano_test

import (
	matasano "github.com/kjaleshire/matasano-go"
	"testing"
)

// Challenge 1 test
func TestHexCharsToValuesBase64(t *testing.T) {
	b64 := matasano.HexCharsToValuesBase64(C1HexString)
	if b64 != C1Base64String {
		t.Errorf("Wrong answer, got %s", b64)
	}
}

// Challenge 2 test
func TestStringXor(t *testing.T) {
	hexString := matasano.StringXor(C2RawString1, C2RawString2)
	if hexString != C2XorResult {
		t.Errorf("Wrong answer, got %s", hexString)
	}
}

// Challenge 3 test
func TestBreakingSingleStringByteKey(t *testing.T) {
	valueBytes := matasano.HexCharsToValues(C3EncodedString)
	decode := matasano.BreakSingleLineByteKey(valueBytes)
	if decode.String != C3DecodedString || decode.ByteKey != C3ByteKey {
		t.Errorf("Wrong answer, got %s with cipher 0x%x", decode.String, decode.ByteKey)
	}
}

// Challenge 4 test
func TestBreakingMultiStringFileByteCipher(t *testing.T) {
	decode := matasano.BreakMultiLineFileByteKey(C4FilePath)
	if decode.String != C4DecodedString || decode.ByteKey != C4ByteKey || decode.Line != C4FileLine {
		t.Errorf("Wrong answer, got %s on line %d with cipher 0x%x", decode.String, decode.Line, decode.ByteKey)
	}
}

// Challenge 5 test
func TestRepeatingKeyXor(t *testing.T) {
	encodedString := matasano.RepeatingKeyXor([]byte(C5UnencodedString), []byte(C5RepeatingKey))
	if encodedString != C5XorResult {
		t.Errorf("Wrong answer, got %s", encodedString)
	}
}

// Challenge 6 tests
func TestHammingDistance(t *testing.T) {
	distance := matasano.HammingBitDistance([]byte(C6HammingString1), []byte(C6HammingString2))
	if distance != C6HammingDistance {
		t.Errorf("Wrong answer, got %d", distance)
	}
}

func TestBreakingRepeatingKeyXorFile(t *testing.T) {
	decode := matasano.BreakRepeatingKeyXorFile(C6FilePath)
	if decode.Key != C6Key {
		t.Errorf("Wrong answer, got key %s\nSize: %d", decode.Key, decode.KeySize)
	}
}

// Challenge 7 tests
func TestDecodingAes128EcbWithKey(t *testing.T) {
	result := matasano.DecodeAes128EcbFile(C7FilePath, []byte(C7Key))
	if result[0:len(C7DecodedFirstLine)] != C7DecodedFirstLine {
		t.Errorf("Wrong answer, got %s", result[0:len(C7DecodedFirstLine)])
	}
}

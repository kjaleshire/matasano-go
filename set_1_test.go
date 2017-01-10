package matasano_test

import (
	"testing"

	matasano "github.com/kjaleshire/matasano-go"
)

// Challenge 1 test
func TestHexCharsToValuesBase64(t *testing.T) {
	result := matasano.HexDecodeBase64(C1HexString)
	if result != C1Base64String {
		t.Errorf("Wrong answer, got %s\n", result)
	}
}

// Challenge 2 test
func TestStringXor(t *testing.T) {
	result := matasano.StringXor(C2RawString1, C2RawString2)
	if result != C2XorResult {
		t.Errorf("Wrong answer, got %s\n", result)
	}
}

// Challenge 3 test
func TestBreakingSingleStringByteKey(t *testing.T) {
	valueBytes := matasano.HexDecodeString(C3EncodedString)
	decodeResult := matasano.BreakSingleLineByteKey(valueBytes)
	if decodeResult.String != C3DecodedString || decodeResult.Key != C3Key {
		t.Errorf("Wrong answer, got %s with key 0x%x\n", decodeResult.String, decodeResult.Key)
	}
}

// Challenge 4 test
func TestBreakingMultiStringFileByteKey(t *testing.T) {
	decodeResult := matasano.BreakMultiLineFileByteKey(C4FilePath)
	if decodeResult.String != C4DecodedString || decodeResult.Key != C4Key {
		t.Errorf("Wrong answer, got %s on line %d with key 0x%x\n",
			decodeResult.String,
			decodeResult.Line,
			decodeResult.Key)
	}
}

// Challenge 5 test
func TestRepeatingKeyXor(t *testing.T) {
	result := matasano.RepeatingKeyXor([]byte(C5UnencodedString), []byte(C5RepeatingKey))
	if result != C5XorResult {
		t.Errorf("Wrong answer, got %s\n", result)
	}
}

// Challenge 6 tests
func TestHammingDistance(t *testing.T) {
	distance := matasano.HammingBitDistance([]byte(C6HammingString1), []byte(C6HammingString2))
	if distance != C6HammingDistance {
		t.Errorf("Wrong answer, got %d\n", distance)
	}
}

func TestBreakingRepeatingKeyXorFile(t *testing.T) {
	cipherBytes := matasano.Base64DecodeString(matasano.DumpFileBytes(C6FilePath))
	key := matasano.BreakRepeatingKeyXorString(string(cipherBytes))
	if string(key) != C6Key {
		t.Errorf("Wrong answer, got key %s\n", key)
	}
}

// Challenge 7 tests
func TestDecodingAesEcbWithKey(t *testing.T) {
	cipherBytes := matasano.Base64DecodeString(matasano.DumpFileBytes(C7FilePath))
	result := matasano.DecryptAesEcbText(string(cipherBytes), []byte(C7Key))
	if result[0:len(C7DecodedFirstLine)] != C7DecodedFirstLine {
		t.Errorf("Wrong answer, got %s\n", result[0:len(C7DecodedFirstLine)])
	}
}

// Challenge 8 tests
func TestDetectEcbFileLine(t *testing.T) {
	lineNumber := matasano.DetectEcbFileLine(C8FilePath)
	if lineNumber != C8LineNumber {
		t.Errorf("Wrong answer, got line %d\n", lineNumber)
	}
}

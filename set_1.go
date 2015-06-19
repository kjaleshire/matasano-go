package matasano

import (
	"encoding/base64"
	"encoding/hex"
)

type DecodeState struct {
	Score  uint
	Cipher uint8
	Line   uint32
	String string
}

// Challenge 1
func HexCharsToValuesBase64(s string) string {
	b := HexCharsToValues(s)

	return base64.StdEncoding.EncodeToString(b)
}

// Challenge 2
func StringXor(s1 string, s2 string) (string, error) {
	bytes_1, err := hex.DecodeString(s1)
	if err != nil {
		return "", err
	}

	bytes_2, err := hex.DecodeString(s2)
	if err != nil {
		return "", err
	}

	var byte_size int
	if len(bytes_1) < len(bytes_2) {
		byte_size = len(bytes_1)
	} else {
		byte_size = len(bytes_2)
	}

	answer_bytes := make([]byte, byte_size)
	for i := 0; i < byte_size; i++ {
		answer_bytes[i] = bytes_1[i] ^ bytes_2[i]
	}

	return hex.EncodeToString(answer_bytes), nil
}

// Challenge 3
func BreakSingleByteCipher(encodedString string) (state DecodeState) {
	valueBytes := HexCharsToValues(encodedString)

	var cipher byte
	for cipher = 0; cipher < 0xFF; cipher++ {
		decodedBytes := make([]byte, len(valueBytes))

		for i, c := range valueBytes {
			decodedBytes[i] = c ^ cipher
		}

		decodedString := string(decodedBytes)

		if newScore := stringScore(decodedString); newScore > state.Score {
			state = DecodeState{Score: newScore, String: decodedString, Cipher: cipher}
		}
	}
	return
}

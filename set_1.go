package matasano

import (
	"encoding/base64"
	"encoding/hex"
)

type DecodeState struct {
	Score  int32
	Cipher uint8
	Line   uint32
	String string
}

func HexCharsToValuesBase64(s string) (string, error) {
	b, err := HexCharsToValues(&s)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(*b), nil
}

func StringXor(s1 string, s2 string) (string, error) {
	bytes_1, err := hex.DecodeString(s1)
	if err != nil {
		return "", err
	}

	bytes_2, err := hex.DecodeString(s2)
	if err != nil {
		return "", err
	}

	byte_size := 0
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

// func BreakSingleCharacterCipher() DecodeState {

// }

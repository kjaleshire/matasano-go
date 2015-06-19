package matasano

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"os"
)

type DecodeState struct {
	Score  int
	Cipher byte
	Line   int
	String string
}

// Challenge 1
func HexCharsToValuesBase64(s string) string {
	return base64.StdEncoding.EncodeToString(HexCharsToValues(s))
}

// Challenge 2
func StringXor(s1 string, s2 string) string {
	bytes_1, err := hex.DecodeString(s1)
	bytes_2, err2 := hex.DecodeString(s2)
	if err != nil || err2 != nil {
		panic(err.Error())
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

	return hex.EncodeToString(answer_bytes)
}

// Challenge 3
func BreakSingleLineByteCipher(encodedString string) (state DecodeState) {
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

// Challenge 4
func BreakMultiLineFileByteCipher(filePath string) (state DecodeState) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	for scanner.Scan() {
		decode := BreakSingleLineByteCipher(scanner.Text())
		if decode.Score > state.Score {
			decode.Line = lineNumber
			state = decode
		}
		lineNumber++
	}
	return
}

func RepeatingKeyXor(stringToEncode, key string) string {
	encodedBytes := make([]byte, len(stringToEncode))

	for i, b := range []byte(stringToEncode) {
		encodedBytes[i] = b ^ key[i%len(key)]
	}

	return hex.EncodeToString(encodedBytes)
}

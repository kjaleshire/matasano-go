package matasano

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"os"
)

type DecodeState struct {
	Score   float32
	Cipher  byte
	Key     string
	KeySize int
	Line    int
	String  string
}

// Challenge 1
func HexCharsToValuesBase64(s string) string {
	return base64.StdEncoding.EncodeToString(HexCharsToValues(s))
}

// Challenge 2
func StringXor(s1 string, s2 string) string {
	bytes_1, err1 := hex.DecodeString(s1)
	bytes_2, err2 := hex.DecodeString(s2)
	if err1 != nil {
		panic(err1.Error())
	}
	if err2 != nil {
		panic(err2.Error())
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
func BreakSingleLineByteCipher(encodedBytes []byte) (state DecodeState) {
	var cipher byte
	for cipher = 0; cipher < 0xFF; cipher++ {
		decodedBytes := make([]byte, len(encodedBytes))

		for i, c := range encodedBytes {
			decodedBytes[i] = c ^ cipher
		}

		decodedString := string(decodedBytes)

		if newScore := float32(stringScore(decodedString)); newScore > state.Score {
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
		encodedBytes := HexCharsToValues(scanner.Text())
		decode := BreakSingleLineByteCipher(encodedBytes)
		if decode.Score > state.Score {
			decode.Line = lineNumber
			state = decode
		}
		lineNumber++
	}
	return
}

// Challenge 5
func RepeatingKeyXor(stringToEncode, key []byte) string {
	encodedBytes := make([]byte, len(stringToEncode))

	for i, b := range stringToEncode {
		encodedBytes[i] = b ^ key[i%len(key)]
	}

	return hex.EncodeToString(encodedBytes)
}

// Challenge 6
func BreakRepeatingKeyXorFile(filePath string) (state DecodeState) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fileString := ""
	for scanner.Scan() {
		fileString += scanner.Text()
	}
	fileBytes, err := base64.StdEncoding.DecodeString(fileString)
	if err != nil {
		panic(err.Error())
	}

	minKeySize := 2
	maxKeySize := 64
	state.Score = 9000

	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
		distance := 0
		passes := len(fileBytes)/keySize - 1
		for i := 0; i < passes; i++ {
			distance += HammingBitDistance(fileBytes[keySize*i:keySize*(i+1)], fileBytes[keySize*(i+1):keySize*(i+2)])
		}
		avgDistance := float32(distance) / float32(passes)
		normalizedDistance := avgDistance / float32(keySize)
		if normalizedDistance < state.Score {
			state = DecodeState{Score: normalizedDistance, KeySize: keySize}
		}
	}

	numberBlocks := len(fileBytes) / state.KeySize
	if len(fileBytes)%state.KeySize > 0 {
		numberBlocks += 1
	}
	blocks := make([][]byte, state.KeySize)

	for i := 0; i < state.KeySize; i++ {
		blocks[i] = make([]byte, numberBlocks)

		for j := 0; j < numberBlocks; j++ {
			if index := j*state.KeySize + i; index < len(fileBytes) {
				blocks[i][j] = fileBytes[index]
			}
		}
	}

	key := make([]byte, state.KeySize)

	for i, block := range blocks {
		s := BreakSingleLineByteCipher(block)
		key[i] = s.Cipher
	}
	state.Key = string(key)
	state.String = string(HexCharsToValues(RepeatingKeyXor(fileBytes, key)))

	return
}

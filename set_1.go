package matasano

import (
	"bytes"
)

type DecodeState struct {
	Score   float32
	Key     string
	KeySize int
	Line    int
	String  string
}

// Challenge 1
func HexDecodeBase64(s string) string {
	return Base64EncodeString(HexDecodeString(s))
}

// Challenge 2
func StringXor(s1 string, s2 string) string {
	bytes1 := HexDecodeString(s1)
	bytes2 := HexDecodeString(s2)

	return HexEncodeToString(ByteSliceXor(bytes1, bytes2))
}

func ByteSliceXor(bytes1, bytes2 []byte) []byte {
	byteSize := len(bytes1)
	if len(bytes2) < len(bytes1) {
		byteSize = len(bytes2)
	}

	result := make([]byte, byteSize)
	for i := 0; i < byteSize; i++ {
		result[i] = bytes1[i] ^ bytes2[i]
	}

	return result
}

// Challenge 3
func BreakSingleLineByteKey(encodedBytes []byte) (state DecodeState) {
	var key byte
	for key = 0; key < 0xFF; key++ {
		decodedBytes := make([]byte, len(encodedBytes))

		for index, currentByte := range encodedBytes {
			decodedBytes[index] = currentByte ^ key
		}

		decodedString := string(decodedBytes)

		if newScore := float32(stringScore(decodedString)); newScore > state.Score {
			state = DecodeState{Score: newScore, String: decodedString, Key: string(key)}
		}
	}
	return
}

// Challenge 4
func BreakMultiLineFileByteKey(filePath string) (state DecodeState) {
	scanner, file := NewScanner(filePath)
	defer file.Close()

	lineNumber := 1
	for scanner.Scan() {
		encodedBytes := HexDecodeString(scanner.Text())
		decode := BreakSingleLineByteKey(encodedBytes)
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

	for index, currentByte := range stringToEncode {
		encodedBytes[index] = currentByte ^ key[index%len(key)]
	}

	return HexEncodeToString(encodedBytes)
}

// Challenge 6
func BreakRepeatingKeyXorString(cipherText string) (state DecodeState) {
	cipherBytes := []byte(cipherText)
	minKeySize := 2
	maxKeySize := 64
	state.Score = 9000

	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
		distance := 0
		passes := len(cipherBytes)/keySize - 1
		for i := 0; i < passes; i++ {
			distance += HammingBitDistance(cipherBytes[keySize*i:keySize*(i+1)], cipherBytes[keySize*(i+1):keySize*(i+2)])
		}
		normalizedAvgDistance := (float32(distance) / float32(passes)) / float32(keySize)
		if normalizedAvgDistance < state.Score {
			state = DecodeState{Score: normalizedAvgDistance, KeySize: keySize}
		}
	}

	numberBlocks := len(cipherBytes) / state.KeySize
	if len(cipherBytes)%state.KeySize > 0 {
		numberBlocks += 1
	}
	blocks := make([][]byte, state.KeySize)

	for i := 0; i < state.KeySize; i++ {
		blocks[i] = make([]byte, numberBlocks)

		for j := 0; j < numberBlocks; j++ {
			if index := j*state.KeySize + i; index < len(cipherBytes) {
				blocks[i][j] = cipherBytes[index]
			}
		}
	}

	key := make([]byte, state.KeySize)

	for index, block := range blocks {
		singleState := BreakSingleLineByteKey(block)
		key[index] = []byte(singleState.Key)[0]
	}
	state.Key = string(key)
	state.String = string(HexDecodeString(RepeatingKeyXor(cipherBytes, key)))

	return
}

// Challenge 7
func DecryptAesEcbString(cipherText string, key []byte) string {
	block, blockSize := setupAesBlock(key, len(cipherText))

	plainBytes := make([]byte, len(cipherText))
	dst := plainBytes
	src := []byte(cipherText)
	for len(src) > 0 {
		block.Decrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}

	return string(plainBytes)
}

// Challenge 8
func DetectEcbFileLine(filePath string) (lineNumber int) {
	scanner, file := NewScanner(filePath)
	defer file.Close()

	lineNumber += 1
	for scanner.Scan() {
		encodedBytes := HexDecodeString(scanner.Text())

		var currentBlock []byte
		for len(encodedBytes[16:]) > 0 {
			currentBlock = encodedBytes[0:16]
			encodedBytes = encodedBytes[16:]
			if bytes.Contains(encodedBytes, currentBlock) {
				return
			}
		}
		lineNumber++
	}
	return
}

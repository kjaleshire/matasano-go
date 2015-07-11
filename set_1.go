package matasano

import (
	"bytes"
)

type DecodeState struct {
	Score  float32
	Key    byte
	Line   int
	String string
}

type KeyState struct {
	Distance float32
	Size     int
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
			state = DecodeState{Score: newScore, String: decodedString, Key: key}
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
func BreakRepeatingKeyXorString(cipherText string) []byte {
	cipherBytes := []byte(cipherText)
	minKeySize := 2
	maxKeySize := 64
	state := KeyState{Distance: 9000}
	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
		total_distance := 0
		passes := len(cipherBytes)/keySize - 1
		for i := 0; i < passes; i++ {
			total_distance += HammingBitDistance(cipherBytes[keySize*i:keySize*(i+1)],
				cipherBytes[keySize*(i+1):keySize*(i+2)])
		}
		normalizedAvgDistance := float32(total_distance) / float32(passes*keySize)
		if normalizedAvgDistance < state.Distance {
			state = KeyState{Distance: normalizedAvgDistance, Size: keySize}
		}
	}

	numberBlocks := len(cipherBytes) / state.Size

	key := make([]byte, state.Size)

	for i := 0; i < state.Size; i++ {
		block := make([]byte, numberBlocks)

		for j := 0; j < numberBlocks; j++ {
			block[j] = cipherBytes[j*state.Size+i]
		}

		singleState := BreakSingleLineByteKey(block)
		key[i] = singleState.Key
	}

	return key
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

package matasano

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"os"
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
	return base64.StdEncoding.EncodeToString(HexDecodeString(s))
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

		for i, c := range encodedBytes {
			decodedBytes[i] = c ^ key
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
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	for i, b := range stringToEncode {
		encodedBytes[i] = b ^ key[i%len(key)]
	}

	return HexEncodeToString(encodedBytes)
}

// Challenge 6
func BreakRepeatingKeyXorFile(filePath string) (state DecodeState) {
	fileString := dumpFileBytes(filePath)

	srcBytes, err := base64.StdEncoding.DecodeString(fileString)
	if err != nil {
		panic(err.Error())
	}

	minKeySize := 2
	maxKeySize := 64
	state.Score = 9000

	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
		distance := 0
		passes := len(srcBytes)/keySize - 1
		for i := 0; i < passes; i++ {
			distance += HammingBitDistance(srcBytes[keySize*i:keySize*(i+1)], srcBytes[keySize*(i+1):keySize*(i+2)])
		}
		avgDistance := float32(distance) / float32(passes)
		normalizedDistance := avgDistance / float32(keySize)
		if normalizedDistance < state.Score {
			state = DecodeState{Score: normalizedDistance, KeySize: keySize}
		}
	}

	numberBlocks := len(srcBytes) / state.KeySize
	if len(srcBytes)%state.KeySize > 0 {
		numberBlocks += 1
	}
	blocks := make([][]byte, state.KeySize)

	for i := 0; i < state.KeySize; i++ {
		blocks[i] = make([]byte, numberBlocks)

		for j := 0; j < numberBlocks; j++ {
			if index := j*state.KeySize + i; index < len(srcBytes) {
				blocks[i][j] = srcBytes[index]
			}
		}
	}

	key := make([]byte, state.KeySize)

	for i, block := range blocks {
		s := BreakSingleLineByteKey(block)
		key[i] = []byte(s.Key)[0]
	}
	state.Key = string(key)
	state.String = string(HexDecodeString(RepeatingKeyXor(srcBytes, key)))

	return
}

// Challenge 7
func DecodeAes128EcbFile(filePath string, key []byte) string {
	fileString := dumpFileBytes(filePath)

	srcBytes, err := base64.StdEncoding.DecodeString(fileString)
	if err != nil {
		panic(err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	blockSize := block.BlockSize()
	if len(srcBytes)%blockSize != 0 {
		panic("Need a multiple of the blocksize")
	}

	decodedBytes := make([]byte, len(srcBytes))
	dst := decodedBytes
	src := []byte(srcBytes)
	for len(src) > 0 {
		block.Decrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}

	return string(decodedBytes)
}

// Challenge 8
func DetectAesEcbFileLine(filePath string) (lineNumber int) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

package matasano

// Challenge 9
func PKCSNo7Padding(message string, blockSize int) string {
	paddingLen := (len(message)/blockSize + 1) * blockSize % len(message)
	padding := make([]byte, paddingLen)
	for i := 0; i < paddingLen; i++ {
		padding[i] = 0x04
	}
	return message + string(padding)
}

// Challenge 10
func DecryptAesCbcString(cipherBytes, key, iv []byte) string {
	block, blockSize := setupAesBlock(key, len(cipherBytes))

	plainBytes := make([]byte, len(cipherBytes))

	dst := plainBytes
	src := cipherBytes
	for len(src) > 0 {
		block.Decrypt(dst, src[:blockSize])
		for i := 0; i < blockSize; i++ {
			dst[i] = dst[i] ^ iv[i]
		}
		iv = src[:blockSize]
		src = src[blockSize:]
		dst = dst[blockSize:]
	}

	return string(plainBytes)
}

func EncryptAesCbcString(plainText string, key, iv []byte) []byte {
	block, blockSize := setupAesBlock(key, len(plainText))

	cipherBytes := make([]byte, len(plainText))

	dst := cipherBytes
	src := []byte(plainText)
	for len(src) > 0 {
		for i := 0; i < blockSize; i++ {
			src[i] = src[i] ^ iv[i]
		}
		block.Encrypt(dst, src[:blockSize])
		iv = dst[:blockSize]
		src = src[blockSize:]
		dst = dst[blockSize:]
	}

	return cipherBytes
}

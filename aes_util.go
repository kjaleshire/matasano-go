package matasano

import (
	"crypto/aes"
	"crypto/cipher"
)

func setupAesBlock(key []byte, textLen int) (cipher.Block, int) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	if textLen%block.BlockSize() != 0 {
		panic("Need a multiple of the blocksize")
	}
	return block, block.BlockSize()
}

package matasano_test

import (
	matasano "github.com/kjaleshire/matasano-go"
	"reflect"
	"testing"
)

// Challenge 9 test
func TestPKCSNo7Padding(t *testing.T) {
	result := matasano.PKCSNo7Padding(C9Block, C9BlockSize)
	if result != C9PaddedBlock {
		t.Errorf("Wrong answer, got %s\n", result)
	}
}

// Challenge 10 test
func TestAes128CbcEncryption(t *testing.T) {
	cipherText := matasano.DumpFileBytes(C10FilePath)

	cipherBytes := matasano.Base64DecodeString(cipherText)
	decryptedResult := matasano.DecryptAesCbcString([]byte(cipherText), []byte(C10Key), []byte(C10IV))

	cipherBytesRedux := matasano.EncryptAesCbcString(decryptedResult, []byte(C10Key), []byte(C10IV))
	if reflect.DeepEqual(cipherBytes, cipherBytesRedux) {
		t.Errorf("Wrong answer, got %s\n", string(cipherBytesRedux))
	}
}

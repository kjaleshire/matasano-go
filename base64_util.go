package matasano

import "encoding/base64"

func Base64DecodeString(s string) []byte {
	plainBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err.Error())
	}
	return plainBytes
}

func Base64EncodeString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

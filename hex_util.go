package matasano

import "encoding/hex"

func HexDecodeString(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err.Error())
	}

	return b
}

func HexEncodeToString(b []byte) string {
	return hex.EncodeToString(b)
}

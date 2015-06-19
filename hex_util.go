package matasano

import (
	"encoding/hex"
)

func HexCharsToValues(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err.Error())
	}

	return b
}

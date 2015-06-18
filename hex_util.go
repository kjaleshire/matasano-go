package matasano

import (
	"encoding/hex"
)

func HexCharsToValues(s *string) (*[]byte, error) {
	b, err := hex.DecodeString(*s)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

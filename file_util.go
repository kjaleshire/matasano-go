package matasano

import (
	"bufio"
	"os"
)

func DumpFileBytes(filePath string) (fileString string) {
	scanner, file := NewScanner(filePath)
	defer file.Close()

	for scanner.Scan() {
		fileString += scanner.Text()
	}
	return
}

func NewScanner(filePath string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}

	return bufio.NewScanner(file), file
}

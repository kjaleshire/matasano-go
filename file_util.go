package matasano

import (
	"bufio"
	"os"
)

func dumpFileBytes(filePath string) (fileString string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileString += scanner.Text()
	}
	return
}

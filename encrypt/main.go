package encrypt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"treuzedev/geheim/shared"
)

func EncryptFile(filePath string, key string) {
	c := make(chan []byte)
	go readFromFile(filePath, c)
	encryptBytes(c)
	// save file
	return
}

func readFromFile(filePath string, c chan []byte) {
	file, err := os.Open(filePath)
	shared.CheckError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		buffer := make([]byte, 16)
		_, err := reader.Read(buffer)
		switch {
		case err == nil:
			c <- buffer
		case err == io.EOF:
			c <- nil
		default:
			shared.CheckError(err)
		}
	}
}

func encryptBytes(c chan []byte) {
	for flag := true; flag; {
		bytes := <-c
		if bytes != nil {
			fmt.Println(bytes)
		} else {
			flag = false
		}
	}
	return
}

func saveToFile() {
	return
}

package encrypt

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"treuzedev/geheim/shared"
)

// TODO: go routines ??
// TODO: channels ?? deadlock ??
func EncryptFile(filePath string, key string) {
	c1 := make(chan shared.ReadFileChannel)
	c2 := make(chan []byte)
	go readFromDecryptedFile(filePath, c1)
	go encryptBytes(key, c1, c2)
	tmpFile := saveBytesToTmpFile(filePath, c2)
	shared.ReplaceFile(filePath, tmpFile)
}

func readFromDecryptedFile(filePath string, c chan shared.ReadFileChannel) {
	file, err := os.Open(filePath)
	shared.CheckError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	readBufferSize := shared.GetReadDecryptedBufferSize()
	shared.ReadFromFile(filePath, c, reader, readBufferSize)
}

func encryptBytes(keyAsString string, c1 chan shared.ReadFileChannel, c2 chan []byte) {
	cipher := shared.GetCipher(keyAsString)
	for r := <-c1; r.Err == nil; r = <-c1 {
		buffer := make([]byte, len(r.Content))
		cipher.Encrypt(buffer, r.Content)
		c2 <- buffer
	}
	c2 <- nil
	close(c2)
}

func saveBytesToTmpFile(filePath string, c chan []byte) string {
	tmpFilePath := fmt.Sprintf("%v", shared.GenerateRandomFilename())
	file, err := os.Create(tmpFilePath)
	shared.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	signFileWithEncryptSignature(writer)
	for bytes := <-c; bytes != nil; bytes = <-c {
		encoded := hex.EncodeToString(bytes)
		_, err := writer.Write([]byte(encoded))
		shared.CheckError(err)
		writer.Flush()
	}
	return tmpFilePath
}

func signFileWithEncryptSignature(writer *bufio.Writer) {
	encryptSignature := shared.GetEncryptSignature()
	_, err := writer.Write(encryptSignature)
	shared.CheckError(err)
	writer.Flush()
}

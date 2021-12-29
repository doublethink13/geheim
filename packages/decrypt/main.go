package decrypt

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"treuzedev/geheim/packages/shared"
)

func Decrypt(filePath, keyAsString string) {
	c1 := make(chan shared.ReadFileChannel)
	c2 := make(chan []byte)
	go readFromEncryptedFile(filePath, c1)
	go decryptBytes(keyAsString, c1, c2)
	tmpFile := saveBytesToTmpFile(filePath, c2)
	shared.ReplaceFile(filePath, tmpFile)
}

func readFromEncryptedFile(filePath string, c chan shared.ReadFileChannel) {
	file, err := os.Open(filePath)
	shared.CheckError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	readEncryptSignature(reader)
	readBufferSize := shared.GetReadEncryptedBufferSize()
	shared.ReadFromFile(filePath, c, reader, readBufferSize)
}

func readEncryptSignature(reader *bufio.Reader) {
	signatureSize := shared.GetReadDecryptedBufferSize()
	buffer := make([]byte, signatureSize)
	_, err := reader.Read(buffer)
	shared.CheckError(err)
}

func decryptBytes(keyAsString string, c1 chan shared.ReadFileChannel, c2 chan []byte) {
	cipher := shared.GetCipher(keyAsString)
	readDecryptBufferSize := shared.GetReadDecryptedBufferSize()
	for r := <-c1; r.Err == nil; r = <-c1 {
		buffer1 := make([]byte, readDecryptBufferSize)
		_, err := hex.Decode(buffer1, r.Content)
		shared.CheckError(err)
		buffer2 := make([]byte, readDecryptBufferSize)
		cipher.Decrypt(buffer2, buffer1)
		c2 <- buffer2
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
	for b := <-c; b != nil; b = <-c {
		b = bytes.Trim(b, "\x00")
		_, err := writer.Write(b)
		shared.CheckError(err)
		writer.Flush()
	}
	return tmpFilePath
}

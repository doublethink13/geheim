package encrypt

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"
	"treuzedev/geheim/shared"
)

var readBufferSize int = shared.GetReadBufferSize()

// TODO: go routines ??
// TODO: channels ?? deadlock ??
func EncryptFile(filePath string, key string) {
	c1 := make(chan ReadFileChannel)
	c2 := make(chan []byte)
	go readFromFile(filePath, c1)
	go encryptBytes(key, c1, c2)
	tmpFile := saveBytesToTmpFile(filePath, c2)
	replaceUnencryptedFile(filePath, tmpFile)
}

func readFromFile(filePath string, c chan ReadFileChannel) {
	file, err := os.Open(filePath)
	shared.CheckError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		buffer := make([]byte, readBufferSize)
		bytesRead, err := reader.Read(buffer)
		shared.CheckError(err)
		switch {
		case bytesRead == readBufferSize:
			c <- ReadFileChannel{buffer, err}
		case bytesRead != readBufferSize:
			trimmed := bytes.Trim(buffer, "\x00")
			c <- ReadFileChannel{trimmed, err}
			c <- ReadFileChannel{[]byte{}, fmt.Errorf("done")}
			close(c)
			return
		}
	}
}

// TODO: actually encrypt the incoming bytes
// TODO: what is a cipher block?
// TODO: what is gcm?
// TODO: what is nonce?
// TODO: what is gcm seal function
func encryptBytes(keyAsString string, c1 chan ReadFileChannel, c2 chan []byte) {
	gcm, nonce := shared.GetCipherResources(keyAsString)
	for r := <-c1; r.err == nil; r = <-c1 {
		randomizedNonce := shared.RandomizeNonce(nonce)
		cipherText := gcm.Seal(randomizedNonce, randomizedNonce, r.content, nil)
		c2 <- cipherText
	}
	c2 <- nil
	close(c2)
}

func saveBytesToTmpFile(filePath string, c chan []byte) string {
	tmpFilePath := fmt.Sprintf("%v", generateRandomFilename())
	file, err := os.Create(tmpFilePath)
	shared.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	signFileWithEncryptSignature(writer)
	for bytes := <-c; bytes != nil; bytes = <-c {
		_, err := writer.Write([]byte(fmt.Sprintf("%x", bytes)))
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

// TODO: replace original file
func replaceUnencryptedFile(originalFile, tmpFile string) {
	// err = os.Remove(filePath)
	// shared.CheckError(err)
	err := os.Rename(tmpFile, fmt.Sprintf("%v%v", originalFile, tmpFile))
	shared.CheckError(err)
}

func generateRandomFilename() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprintf("%v", r.Intn(10000))
}

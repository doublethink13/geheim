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

var readBufferSize int = 16

func EncryptFile(filePath string, key string) {
	c1 := make(chan ReadFileChannel)
	c2 := make(chan []byte)
	go readFromFile(filePath, c1)
	go encryptBytes(c1, c2)
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
		switch {
		case bytesRead == readBufferSize:
			c <- ReadFileChannel{buffer, err}
		case bytesRead == 0:
			close(c)
			return
		case bytesRead != readBufferSize:
			trimmed := bytes.Trim(buffer, "\x00")
			c <- ReadFileChannel{trimmed, err}
			close(c)
			return
		default:
			close(c)
			shared.CheckError(err)
			return
		}
	}
}

// TODO: actually encrypt the incoming bytes
func encryptBytes(c1 chan ReadFileChannel, c2 chan []byte) {
	for r := <-c1; r.err == nil; r = <-c1 {
		c2 <- r.content
	}
	c2 <- nil
	close(c2)
}

func saveBytesToTmpFile(filePath string, c chan []byte) string {
	tmpFile := fmt.Sprintf("%v", generateRandomFilename())
	file, err := os.Create(tmpFile)
	shared.CheckError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for bytes := <-c; bytes != nil; bytes = <-c {
		_, err := writer.Write(bytes)
		shared.CheckError(err)
		writer.Flush()
	}
	return tmpFile
}

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

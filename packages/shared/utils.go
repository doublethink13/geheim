package shared

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// TODO:
// without peek, reading stops midway without any error
// go panics with "encoding/hex: invalid byte: U+0000"
func ReadFromFile(filePath string, c chan ReadFileChannel, reader *bufio.Reader, readBufferSize int) {
	for {
		reader.Peek(readBufferSize)
		buffer := make([]byte, readBufferSize)
		bytesRead, err := reader.Read(buffer)
		switch {
		case bytesRead == readBufferSize:
			CheckError(err)
			c <- ReadFileChannel{Content: buffer, Err: err}
		case bytesRead != readBufferSize:
			if err != io.EOF {
				CheckError(err)
			}
			c <- ReadFileChannel{Content: buffer, Err: err}
			c <- ReadFileChannel{Content: []byte{}, Err: fmt.Errorf("done")}
			close(c)
			return
		}
	}
}

func CheckIfEncrypted(filePath string) bool {
	file, err := os.Open(filePath)
	CheckError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, readDecryptedBufferSize)
	_, err = reader.Read(buffer)
	CheckError(err)
	encryptSignature := GetEncryptSignature()
	return CompareByteSlices(encryptSignature, buffer)
}

func GenerateRandomFilename() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprintf("%v", r.Intn(10000))
}

func ReplaceFile(originalFile, tmpFile string) {
	err := os.Remove(originalFile)
	CheckError(err)
	err = os.Rename(tmpFile, originalFile)
	CheckError(err)
}

func CompareStringSlices(a, b []string) bool {
	switch {
	case len(a) != len(b):
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareByteSlices(a, b []byte) bool {
	switch {
	case len(a) != len(b):
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

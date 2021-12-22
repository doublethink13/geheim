package shared

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func ReadFromFile(filePath string, c chan ReadFileChannel, reader *bufio.Reader, readBufferSize int) {
	for {
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

// TODO: replace original file
func ReplaceFile(originalFile, tmpFile string) {
	// err = os.Remove(filePath)
	// shared.CheckError(err)
	err := os.Rename(tmpFile, fmt.Sprintf("%v%v", originalFile, tmpFile))
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

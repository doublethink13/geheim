package shared

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
)

func ReadFromFile(filePath string, readFileChannel chan ReadFileChannel, reader *bufio.Reader, readBufferSize int) {
	for {
		// without peek, reading stops midway without any error
		// go panics with "encoding/hex: invalid byte: U+0000"
		_, err := reader.Peek(readBufferSize)
		if !errors.Is(err, io.EOF) {
			CheckError(err, nil)
		}

		buffer := make([]byte, readBufferSize)
		bytesRead, err := reader.Read(buffer)

		switch {
		case bytesRead == readBufferSize:
			CheckError(err, nil)
			readFileChannel <- ReadFileChannel{Content: buffer, Err: err}
		case bytesRead != readBufferSize:
			if !errors.Is(err, io.EOF) {
				CheckError(err, nil)
			}

			readFileChannel <- ReadFileChannel{Content: buffer, Err: err}
			readFileChannel <- ReadFileChannel{Content: []byte{}, Err: fmt.Errorf("done")}

			close(readFileChannel)

			return
		}
	}
}

func CheckIfEncrypted(filePath string) (isEncrypted bool) {
	file, err := os.Open(filePath)
	CheckError(err, nil)

	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, readDecryptedBufferSize)

	_, err = reader.Read(buffer)
	CheckError(err, nil)

	encryptSignature := GetEncryptSignature()

	return CompareByteSlices(encryptSignature, buffer)
}

func GenerateRandomFilename() (filename string) {
	randomNumberMax := 1000
	n, err := rand.Int(rand.Reader, big.NewInt(int64(randomNumberMax)))
	CheckError(err, nil)

	return fmt.Sprintf("%v", n)
}

func ReplaceFile(originalFile, tmpFile string) {
	err := os.Remove(originalFile)
	CheckError(err, nil)

	err = os.Rename(tmpFile, originalFile)
	CheckError(err, nil)
}

func CompareStringSlices(sliceA, sliceB []string) (areEqual bool) {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i := 0; i < len(sliceA); i++ {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}

func CompareByteSlices(sliceA, sliceB []byte) (areEqual bool) {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i := 0; i < len(sliceA); i++ {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}

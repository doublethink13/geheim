package shared

import (
	"bufio"
	"os"
)

func CheckIfEncrypted(filePath string) bool {
	file, err := os.Open(filePath)
	CheckError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, readBufferSize)
	_, err = reader.Read(buffer)
	CheckError(err)
	encryptSignature := GetEncryptSignature()
	return CompareByteSlices(encryptSignature, buffer)
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

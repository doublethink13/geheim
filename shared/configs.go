package shared

import "fmt"

var readBufferSize int = 16
var encryptSignature = []byte{100, 111, 100, 113, 122, 111, 117, 100, 101, 97, 104, 100, 105, 100, 106, 108}
var encrytionKeySize int = 16
var bytesToAppend = [16]byte{98, 109, 102, 115, 124, 113, 115, 98, 99, 99, 106, 98, 107, 102, 108, 106}

func GetReadBufferSize() int {
	return readBufferSize
}

func GetEncryptSignature() []byte {
	if len(encryptSignature) != readBufferSize {
		panic(fmt.Errorf("encryptSignature needs to be of size %v", readBufferSize))
	}
	return encryptSignature
}

func GetEncryptionKeySize() int {
	return encrytionKeySize
}

func GetBytesToAppend() [16]byte {
	return bytesToAppend
}

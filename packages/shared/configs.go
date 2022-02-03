package shared

import "fmt"

const (
	readDecryptedBufferSize int = 16
	readEncryptedBufferSize int = 32
	encrytionKeySize        int = 16
)

func GetReadDecryptedBufferSize() (size int) {
	return readDecryptedBufferSize
}

func GetReadEncryptedBufferSize() (size int) {
	return readEncryptedBufferSize
}

func GetEncryptSignature() (signature []byte) {
	encryptSignature := []byte{100, 111, 100, 113, 122, 111, 117, 100, 101, 97, 104, 100, 105, 100, 106, 108}

	if len(encryptSignature) != readDecryptedBufferSize {
		panic(fmt.Errorf("encryptSignature needs to be of size : %v", readDecryptedBufferSize))
	}

	return encryptSignature
}

func GetEncryptionKeySize() (size int) {
	return encrytionKeySize
}

func GetBytesToAppend() (bytes [16]byte) {
	return [16]byte{98, 109, 102, 115, 124, 113, 115, 98, 99, 99, 106, 98, 107, 102, 108, 106}
}

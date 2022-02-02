package encrypt

//nolint
import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"treuzedev/geheim/packages/logging"
	"treuzedev/geheim/packages/shared"
)

func File(filePath string, key string) {
	logger := logging.GetLogger()

	readFromFileChannel := make(chan shared.ReadFileChannel)
	encryptBytesChannel := make(chan []byte)

	go readFromDecryptedFile(filePath, readFromFileChannel)
	go encryptBytes(key, readFromFileChannel, encryptBytesChannel)

	tmpFile := saveBytesToTmpFile(filePath, encryptBytesChannel)
	shared.ReplaceFile(filePath, tmpFile)

	logger.Log(
		logging.Info,
		logging.InfoLogLevel,
		fmt.Sprintf("Encrypted file: %v", filePath),
	)
}

func readFromDecryptedFile(filePath string, readFromFileChannel chan shared.ReadFileChannel) {
	file, err := os.Open(filePath)
	shared.CheckError(err, nil)

	defer file.Close()

	reader := bufio.NewReader(file)
	readBufferSize := shared.GetReadDecryptedBufferSize()
	shared.ReadFromFile(
		filePath,
		readFromFileChannel,
		reader,
		readBufferSize,
	)
}

func encryptBytes(
	keyAsString string,
	readFromFileChannel chan shared.ReadFileChannel,
	encryptBytesChannel chan []byte,
) {
	cipher := shared.GetCipher(keyAsString)

	for r := <-readFromFileChannel; r.Err == nil; r = <-readFromFileChannel {
		buffer := make([]byte, len(r.Content))
		cipher.Encrypt(buffer, r.Content)
		encryptBytesChannel <- buffer
	}

	encryptBytesChannel <- nil
	close(encryptBytesChannel)
}

func saveBytesToTmpFile(filePath string, encryptBytesChannel chan []byte) (tmpFilePath string) {
	tmpFilePath = fmt.Sprintf("%v", shared.GenerateRandomFilename())
	file, err := os.Create(tmpFilePath)
	shared.CheckError(err, &filePath)

	defer file.Close()

	writer := bufio.NewWriter(file)
	signFileWithEncryptSignature(writer)

	for bytes := <-encryptBytesChannel; bytes != nil; bytes = <-encryptBytesChannel {
		encoded := hex.EncodeToString(bytes)
		_, err := writer.Write([]byte(encoded))
		shared.CheckError(err, &filePath)

		err = writer.Flush()
		shared.CheckError(err, &filePath)
	}

	return tmpFilePath
}

func signFileWithEncryptSignature(writer *bufio.Writer) {
	encryptSignature := shared.GetEncryptSignature()
	_, err := writer.Write(encryptSignature)
	shared.CheckError(err, nil)
	writer.Flush()
}

package decrypt

//nolint
import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"treuzedev/geheim/packages/logging"
	"treuzedev/geheim/packages/shared"
)

func Decrypt(filePath, keyAsString string) {
	logger := logging.GetLogger()

	readFromFileChannel := make(chan shared.ReadFileChannel)
	decryptBytesChannel := make(chan []byte)

	go readFromEncryptedFile(filePath, readFromFileChannel)
	go decryptBytes(keyAsString, readFromFileChannel, decryptBytesChannel)

	tmpFile := saveBytesToTmpFile(filePath, decryptBytesChannel)
	shared.ReplaceFile(filePath, tmpFile)

	logger.Log(
		logging.Info,
		logging.InfoLogLevel,
		fmt.Sprintf("Decrypted file: %v", filePath),
	)
}

func readFromEncryptedFile(filePath string, readFromFileChannel chan shared.ReadFileChannel) {
	file, err := os.Open(filePath)
	shared.CheckError(err, nil)

	defer file.Close()

	reader := bufio.NewReader(file)
	readEncryptSignature(reader)

	readBufferSize := shared.GetReadEncryptedBufferSize()

	shared.ReadFromFile(
		filePath,
		readFromFileChannel,
		reader,
		readBufferSize,
	)
}

func readEncryptSignature(reader io.Reader) {
	signatureSize := shared.GetReadDecryptedBufferSize()
	buffer := make([]byte, signatureSize)
	_, err := reader.Read(buffer)
	shared.CheckError(err, nil)
}

func decryptBytes(
	keyAsString string,
	readFromFileChannel chan shared.ReadFileChannel,
	decryptBytesChannel chan []byte,
) {
	cipher := shared.GetCipher(keyAsString)
	readDecryptBufferSize := shared.GetReadDecryptedBufferSize()

	for r := <-readFromFileChannel; r.Err == nil; r = <-readFromFileChannel {
		buffer1 := make([]byte, readDecryptBufferSize)
		_, err := hex.Decode(buffer1, r.Content)
		shared.CheckError(err, nil)

		buffer2 := make([]byte, readDecryptBufferSize)
		cipher.Decrypt(buffer2, buffer1)
		decryptBytesChannel <- buffer2
	}

	decryptBytesChannel <- nil
	close(decryptBytesChannel)
}

func saveBytesToTmpFile(filePath string, decryptBytesChannel chan []byte) (tmpFilePath string) {
	tmpFilePath = fmt.Sprintf("%v", shared.GenerateRandomFilename())
	file, err := os.Create(tmpFilePath)
	shared.CheckError(err, &filePath)

	defer file.Close()

	writer := bufio.NewWriter(file)

	for b := <-decryptBytesChannel; b != nil; b = <-decryptBytesChannel {
		b = bytes.Trim(b, "\x00")
		_, err := writer.Write(b)
		shared.CheckError(err, &filePath)

		writer.Flush()
	}

	return tmpFilePath
}

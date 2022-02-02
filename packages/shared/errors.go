package shared

import (
	"errors"
	"fmt"
)

type encryptSignatureSizeError struct {
	size int
}

var errFoo = errors.New("encryptSignature needs to be of size")

func (err *encryptSignatureSizeError) Error() error {
	return fmt.Errorf("%w : %v", errFoo, err.size)
}

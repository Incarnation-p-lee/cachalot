package utils

import (
	"fmt"
	"os"
	"errors"
)

func CloseFile(file *os.File) error {
	if file == nil {
		return errors.New("Cannot close ni file")
	}

	if err := file.Close(); err != nil {
		return errors.New(fmt.Sprintf("Cannot close file %+v due to %+v\n", file, err))
	}

	return nil
}


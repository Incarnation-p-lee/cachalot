package utils

import (
	"errors"
	"fmt"
	"os"
)

func CloseFile(file *os.File) error {
	if file == nil {
		return errors.New("Cannot close ni file")
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("Cannot close file %+v due to %+v\n", file, err)
	}

	return nil
}

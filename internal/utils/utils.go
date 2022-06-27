package utils

import (
	"errors"
	"fmt"
	"os"
)

// CloseFile will close the file of os.File, with error returned if any.
func CloseFile(file *os.File) error {
	if file == nil {
		return errors.New("cannot close ni file")
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("cannot close file %+v due to %+v", file, err)
	}

	return nil
}

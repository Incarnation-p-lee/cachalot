package utils

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
	"os"
)

func TestCloseFile(t *testing.T) {
	file, _ := os.Open("/proc/meminfo")

	assert.IsNotNil(t, CloseFile(nil), "close nil file should have error")
	assert.IsNil(t, CloseFile(file), "close normal file should have nil error")
}

func TestCloseFileTwice(t *testing.T) {
	file, _ := os.Open("/proc/meminfo")
	file.Close()

	assert.IsNotNil(t, CloseFile(file), "close file twice should have error")
}


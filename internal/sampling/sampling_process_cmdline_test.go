package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
)

func TestSampleCmdLine(t *testing.T) {
	cmdLineChan := make(chan string)
	defer close(cmdLineChan)

	go sampleCmdLine(1, cmdLineChan)
	cmdLine := <-cmdLineChan

	assert.IsTrue(t, len(cmdLine) > 0, "cmdLine length should be greater than 0.")
}

func TestSampleCmdLineInvalidPID(t *testing.T) {
	cmdLineChan := make(chan string)
	defer close(cmdLineChan)

	go sampleCmdLine(100000000, cmdLineChan)
	cmdLine := <-cmdLineChan

	assert.IsEqual(t, unknownCmdLine, cmdLine, "cmdLine should be unknown.")
}

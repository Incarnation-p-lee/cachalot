package sampling

import (
    "testing"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestSampleCmdLine(t *testing.T) {
    testPID := 1
    cmdLine := sampleCmdLine(testPID)

    assert.IsTrue(t, len(cmdLine) > 0, "cmdLine length should be greater than 0.")
}

func TestSampleCmdLineInvalidPID(t *testing.T) {
    testPID := 10000000
    cmdLine := sampleCmdLine(testPID)

    assert.IsEqual(t, unknownCmdLine, cmdLine, "cmdLine should be unknown.")
}


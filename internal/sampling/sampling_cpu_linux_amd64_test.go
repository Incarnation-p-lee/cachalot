package sampling

import (
    "testing"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestSampleCPU(t *testing.T) {
    testPID := 1

    cpuStat := sampleCPU(testPID)

    assert.IsTrue(t, cpuStat.JiffiesInTotal != invalidJiffies,
        "cpu total jiffies should not be invalid.")
    assert.IsTrue(t, cpuStat.JiffiesInTotal != defaultJiffies,
        "cpu total jiffies should not be default value.")

    assert.IsTrue(t, cpuStat.JiffiesUsed != invalidJiffies,
        "cpu process used jiffies should not be invalid.")
    assert.IsTrue(t, cpuStat.JiffiesUsed != defaultJiffies,
        "cpu process used jiffies should not be default value.")
}


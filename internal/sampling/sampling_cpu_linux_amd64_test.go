package sampling

import (
    "testing"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestSampleCPUStat(t *testing.T) {
    testPID := 1
    cpuStat := sampleCPUStat(testPID)

    assert.IsTrue(t, cpuStat.JiffiesInTotal != invalidJiffies,
        "cpu total jiffies should not be invalid.")
    assert.IsTrue(t, cpuStat.JiffiesInTotal != defaultJiffies,
        "cpu total jiffies should not be default value.")

    assert.IsTrue(t, cpuStat.UsageInPercentage <= 100.0,
        "the cpu usage percentage should be less than or equal to 100.0")
}

func TestSampleCPUJiffies(t *testing.T) {
    testPID := 1
    totalJiffies, processJiffies := sampleCPUJiffies(testPID)

    assert.IsTrue(t, processJiffies <= totalJiffies,
        "the process jiffies should be less than or eual to total jiffies")

    assert.IsTrue(t, processJiffies != invalidJiffies,
        "the process used jiffies should not be invalid.")

    assert.IsTrue(t, processJiffies != defaultJiffies,
        "the process used jiffies should not be default value.")
}

func TestTestSampleCPUJiffiesInvalidPID(t *testing.T) {
    testPID := 10000000
    totalJiffies, processJiffies := sampleCPUJiffies(testPID)

    assert.IsTrue(t, processJiffies <= totalJiffies,
        "the process jiffies should be less than or eual to total jiffies")

    assert.IsTrue(t, processJiffies == invalidJiffies,
        "the invalid process id used jiffies should be invalid.")

    assert.IsTrue(t, processJiffies != defaultJiffies,
        "the invalid process id used jiffies should not be default value.")
}


package sampling

import (
    "testing"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestSampleCPU(t *testing.T) {
    testPId := 1

    cpuStat := sampleCPU(testPId)

    assert.IsEqual(t, cpuStat.JiffiesUsed, invalidJiffies, "process jiffies should be invalid")
    assert.IsTrue(t, cpuStat.JiffiesInTotal != invalidJiffies,
        "cpu total jiffies should not be invalid.")
}


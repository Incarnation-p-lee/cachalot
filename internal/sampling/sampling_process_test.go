package sampling

import (
    "testing"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestGetAllProcessID(t *testing.T) {
    allPIDs := getAllProcessID()

    assert.IsTrue(t, len(allPIDs) > 0, "All proccess ID count should not be 0.")
}

func TestSampleAllProcess(t *testing.T) {
    ops := options.CreateOptions()
    processes := sampleAllProcess(ops)

    assert.IsTrue(t, len(processes) > 0, "All proccess slice count should not be 0.")

    for i := 0; i < len(processes) - 1; i++ {
        a, b := processes[i], processes[i + 1]

        assert.IsTrue(t, a.CPU.UsageInPercentage >= b.CPU.UsageInPercentage,
            "the process usage should be sorted in desc order")
    }
}


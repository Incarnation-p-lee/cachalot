package sampling

import (
    "testing"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func isContains(processIDs []int, pID int) bool {
    for _, processID := range processIDs {
        if processID == pID {
            return true
        }
    }

    return false
}

func TestGetAllProcessID(t *testing.T) {
    allPIDs := getAllProcessID()

    assert.IsTrue(t, len(allPIDs) > 0, "all proccess ID count should not be 0.")
    assert.IsTrue(t, isContains(allPIDs, 1), "all proccess ID should contain pID 1.")
}

func TestSampleAllProcess(t *testing.T) {
    ops := options.CreateOptions()
    processes := sampleAllProcess(ops)

    assert.IsTrue(t, len(processes) > 0, "all proccess slice count should not be 0.")

    for i := 0; i < len(processes) - 1; i++ {
        a, b := processes[i], processes[i + 1]

        assert.IsTrue(t, a.CPUStat.UsageInPercentage >= b.CPUStat.UsageInPercentage,
            "the process usage should be sorted in desc order")
    }
}

func TestSampleOneProcessSnapshotNilOptions(t *testing.T) {
    testPIDChan, testProcessChan := make(chan int, 1), make(chan snapshot.Process, 1)

    go sampleOneProcessSnapshot(nil, testPIDChan, testProcessChan)

    defer close(testPIDChan)
    defer close(testProcessChan)

    testProcess := <- testProcessChan

    assert.IsEqual(t, 0, testProcess.PID, "nil options will have 0 pID for process")
    assert.IsEqual(t, 0, len(testProcess.CmdLine), "nil options will have empty cmdLine for process")
    assert.IsEqual(t, 0, testProcess.CPUStat.JiffiesUsed,
        "nil options will have 0 jiffies used for process")
    assert.IsEqual(t, 0, testProcess.CPUStat.JiffiesInTotal,
        "nil options will have 0 jiffies in total for process")
}

func TestSampleOneProcessSnapshot(t *testing.T) {
    ops := options.CreateOptions()
    testPIDChan, testProcessChan := make(chan int, 1), make(chan snapshot.Process, 1)

    go sampleOneProcessSnapshot(ops, testPIDChan, testProcessChan)

    defer close(testPIDChan)
    defer close(testProcessChan)

    testPID := 1
    testPIDChan <- testPID
    testProcess := <- testProcessChan

    assert.IsEqual(t, testPID, testProcess.PID, "process will have the same pID")
    assert.IsTrue(t, len(testProcess.CmdLine) > 0, "process will have text for cmdLine")
    assert.IsTrue(t, testProcess.CPUStat.JiffiesUsed >= 0,
        "process will have 0 jiffies used for process")
    assert.IsTrue(t, testProcess.CPUStat.JiffiesInTotal > 0,
        "process will have positive jiffies in total")
}


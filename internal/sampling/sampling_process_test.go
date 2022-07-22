package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/options"
	"testing"
)

func isContains(processIDs []int, pID int) bool {
	for _, processID := range processIDs {
		if processID == pID {
			return true
		}
	}

	return false
}

func TestGetAllProcessIDs(t *testing.T) {
	allPIDs := getAllProcessIDs()

	assert.IsTrue(t, len(allPIDs) > 0, "all proccess ID count should not be 0.")
	assert.IsTrue(t, isContains(allPIDs, 1), "all proccess ID should contain pID 1.")
}

func TestSampleAllProcess(t *testing.T) {
	ops := options.CreateOptions()
	ops.AppendOption(options.Option{
		Key: options.ProcessIDs,
		Val: options.AllProcessIDs,
	})

	processes := sampleProcesses(ops, snapshot.Snapshot{})

	assert.IsTrue(t, len(processes) > 0, "all proccess slice count should not be 0.")
}

func TestSampleOneProcessSnapshotNilOptions(t *testing.T) {
	testProcessChan := make(chan snapshot.Process, 1)
	defer close(testProcessChan)

	go sampleOneProcessSnapshot(nil, 1, snapshot.Snapshot{}, testProcessChan)
	testProcess := <-testProcessChan

	assert.IsEqual(t, 0, testProcess.PID, "nil options will have 0 pID for process")
	assert.IsEqual(t, 0, len(testProcess.CmdLine),
		"nil options will have empty cmdLine for process")
	assert.IsEqual(t, 0, testProcess.CPUStat.JiffiesUsed,
		"nil options will have 0 jiffies used for process")
	assert.IsEqual(t, 0, testProcess.CPUStat.JiffiesInTotal,
		"nil options will have 0 jiffies in total for process")
}

func TestSampleOneProcessSnapshot(t *testing.T) {
	testPID := 1
	ops := options.CreateOptions()

	testProcessChan := make(chan snapshot.Process, 1)
	defer close(testProcessChan)

	go sampleOneProcessSnapshot(ops, testPID, snapshot.Snapshot{}, testProcessChan)
	testProcess := <-testProcessChan

	assert.IsEqual(t, testPID, testProcess.PID, "process will have the same pID")
	assert.IsTrue(t, len(testProcess.CmdLine) > 0, "process will have text for cmdLine")
	assert.IsTrue(t, testProcess.CPUStat.JiffiesUsed >= 0,
		"process will have 0 jiffies used for process")
	assert.IsTrue(t, testProcess.CPUStat.JiffiesInTotal > 0,
		"process will have positive jiffies in total")
}

func TestGetProcessIDs(t *testing.T) {
	testPIDs := "1,2,3"
	processIDs := getProcessIDs(testPIDs)

	assert.IsEqual(t, 3, len(processIDs), "there should be 3 process IDs")
	assert.IsEqual(t, 1, processIDs[0], "first process ID should be 1")
	assert.IsEqual(t, 2, processIDs[1], "second process ID should be 2")
	assert.IsEqual(t, 3, processIDs[2], "third process ID should be 3")
}

func TestGetProcessIDsInvalid(t *testing.T) {
	testPIDs := "1,b,2,a,3"
	processIDs := getProcessIDs(testPIDs)

	assert.IsEqual(t, 3, len(processIDs), "there should be 3 process IDs")
	assert.IsEqual(t, 1, processIDs[0], "first process ID should be 1")
	assert.IsEqual(t, 2, processIDs[1], "second process ID should be 2")
	assert.IsEqual(t, 3, processIDs[2], "third process ID should be 3")
}

package snapshot

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
	"time"
)

func createProcess(pID int, cmdLine string, cpuStat CPUStat) Process {
	return Process{
		PID:     pID,
		CmdLine: cmdLine,
		CPUStat: cpuStat,
	}
}

func TestCreateCPUStat(t *testing.T) {
	used, limited := 250, 500
	stat := CreateCPUStat(used, limited)

	assert.IsEqual(t, used, stat.JiffiesUsed, "cpu stat should have the same used value.")
	assert.IsEqual(t, limited, stat.JiffiesInTotal, "cpu stat should have the same limited value.")
	assert.IsEqual(t, float64(used*100/limited), stat.UsageInPercentage,
		"cpu stat should have the same percentage value.")
}

func TestSetUsage(t *testing.T) {
	process := createProcess(123, "ls -l", CPUStat{})

	assert.IsEqual(t, 0.0, process.CPUStat.UsageInPercentage, "process cpu stat should be zero.")

	used, limited := 250, 500
	stat := CreateCPUStat(used, limited)

	process.SetCPUStat(stat)

	assert.IsEqual(t, float64(used*100/limited), process.CPUStat.UsageInPercentage,
		"cpu stat should have the same percentage value.")
}

func TestAppendProcess(t *testing.T) {
	process := createProcess(123, "ls -l", CPUStat{})
	snapshot := Snapshot{
		Timestamp: time.Now(),
		Processes: []Process{},
	}

	assert.IsEqual(t, 0, len(snapshot.Processes), "snapshot processes count should be 0.")

	snapshot.AppendProcess(process)

	assert.IsEqual(t, 1, len(snapshot.Processes), "snapshot processes count should be 0.")
}

func TestAppendProcesses(t *testing.T) {
	snapshot := Snapshot{
		Timestamp: time.Now(),
		Processes: []Process{
			Process{},
		},
	}
	processes := []Process{
		createProcess(123, "ls -l", CPUStat{}),
		createProcess(123, "ls -l", CPUStat{}),
	}

	snapshot.AppendProcesses(processes)

	assert.IsEqual(t, 3, len(snapshot.Processes), "processes size should be 3")
}

func TestAppendProcessesWithExistNil(t *testing.T) {
	snapshot := Snapshot{
		Timestamp: time.Now(),
		Processes: nil,
	}
	processes := []Process{
		createProcess(123, "ls -l", CPUStat{}),
		createProcess(123, "ls -l", CPUStat{}),
	}

	snapshot.AppendProcesses(processes)

	assert.IsEqual(t, 2, len(snapshot.Processes), "processes size should be 2")
}

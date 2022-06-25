package snapshot

import (
    "time"
    "testing"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func createProcess(pID int, cmdLine string, cpuStat CPUStat) Process {
	return Process {
		PID: pID,
		CmdLine: cmdLine,
		CPUStat: cpuStat,
	}
}

func TestCreateSnapshot(t *testing.T) {
    testLastHour, processes := time.Now().Add(-time.Hour), []Process {}
    snapshot := CreateSnapshot(testLastHour, processes)

    assert.IsEqual(t, testLastHour, snapshot.Timestamp, "snapshot should have the same timestamp.")
    assert.IsEqual(t, 0, len(snapshot.Processes), "snapshot should have empty processes slice.")
}

func TestCreateCPUStat(t *testing.T) {
    used, limited := 250, 500
    stat := CreateCPUStat(used, limited)

    assert.IsEqual(t, used, stat.JiffiesUsed, "cpu stat should have the same used value.") 
    assert.IsEqual(t, limited, stat.JiffiesInTotal, "cpu stat should have the same limited value.") 
    assert.IsEqual(t, float64(used * 100 / limited), stat.UsageInPercentage,
        "cpu stat should have the same percentage value.") 
}

func TestSetUsage(t *testing.T) {
    process := createProcess(123, "ls -l", CPUStat {})

    assert.IsEqual(t, 0.0, process.CPUStat.UsageInPercentage, "process cpu stat should be zero.")

    used, limited := 250, 500
    stat := CreateCPUStat(used, limited)

    process.SetCPUStat(stat)

    assert.IsEqual(t, float64(used * 100 / limited), process.CPUStat.UsageInPercentage,
        "cpu stat should have the same percentage value.") 
}

func TestAppendProcess(t *testing.T) {
    process := createProcess(123, "ls -l", CPUStat {})
    snapshot := CreateSnapshot(time.Now(), []Process {})

    assert.IsEqual(t, 0, len(snapshot.Processes), "snapshot processes count should be 0.")

    snapshot.AppendProcess(process)

    assert.IsEqual(t, 1, len(snapshot.Processes), "snapshot processes count should be 0.")
}


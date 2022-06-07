package snapshot

import (
    "time"
    "testing"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestCreateSnapshot(t *testing.T) {
    testLastHour := time.Now().Add(-time.Hour)
    snapshot := CreateSnapshot(testLastHour)

    assert.IsEqual(t, testLastHour, snapshot.Timestamp, "snapshot should have the same timestamp.")
}

func TestCreateProcess(t *testing.T) {
    cmdLine, pID := "bash -c", 123
    process := CreateProcess(cmdLine, pID)

    assert.IsEqual(t, cmdLine, process.CmdLine, "process should have the same cmd line.")
    assert.IsEqual(t, pID, process.PID, "process should have the same id.")
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
    process := CreateProcess("ls -l", 123)

    assert.IsEqual(t, 0.0, process.CPU.UsageInPercentage, "process cpu stat should be zero.")

    used, limited := 250, 500
    stat := CreateCPUStat(used, limited)

    process.SetCPUStat(stat)

    assert.IsEqual(t, float64(used * 100 / limited), process.CPU.UsageInPercentage,
        "cpu stat should have the same percentage value.") 
}

func TestAppendProcess(t *testing.T) {
    process := CreateProcess("ls -l", 123)
    snapshot := CreateSnapshot(time.Now())

    assert.IsEqual(t, 0, len(snapshot.Processes), "snapshot processes count should be 0.")

    snapshot.AppendProcess(process)

    assert.IsEqual(t, 1, len(snapshot.Processes), "snapshot processes count should be 0.")
}


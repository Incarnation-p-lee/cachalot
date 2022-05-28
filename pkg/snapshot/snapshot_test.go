package snapshot

import (
    "time"
    "testing"
    "pkg/assert"
)

func TestCreateSnapshot(t *testing.T) {
    testLastHour := time.Now().Add(-time.Hour)
    snapshot := CreateSnapshot(testLastHour)

    assert.IsEqual(t, testLastHour, snapshot.Timestamp, "snapshot should have the same timestamp.")
}

func TestCreateProcess(t *testing.T) {
    cmdLine := "bash -c"
    process := CreateProcess(cmdLine)

    assert.IsEqual(t, cmdLine, process.CmdLine, "process should have the same cmd line.")
}

func TestCreateCpuStat(t *testing.T) {
    used, limited := 250.0, 500.0
    stat := CreateCpuStat(used, limited)

    assert.IsEqual(t, used, stat.MCoreUsed, "cpu stat should have the same used value.") 
    assert.IsEqual(t, limited, stat.MCoreLimited, "cpu stat should have the same limited value.") 
    assert.IsEqual(t, used * 100.0 / limited, stat.UsageInPercentage,
        "cpu stat should have the same percentage value.") 
}

func TestSetUsage(t *testing.T) {
    process := CreateProcess("ls -l")

    assert.IsEqual(t, 0.0, process.Cpu.UsageInPercentage, "process cpu stat should be zero.")

    used, limited := 250.0, 500.0
    stat := CreateCpuStat(used, limited)

    process.SetCpuStat(stat)

    assert.IsEqual(t, used * 100.0 / limited, process.Cpu.UsageInPercentage,
        "cpu stat should have the same percentage value.") 
}

func TestAppendProcess(t *testing.T) {
    process := CreateProcess("ls -l")
    snapshot := CreateSnapshot(time.Now())

    assert.IsEqual(t, 0, len(snapshot.Processes), "snapshot processes count should be 0.")

    snapshot.AppendProcess(process)

    assert.IsEqual(t, 1, len(snapshot.Processes), "snapshot processes count should be 0.")
}


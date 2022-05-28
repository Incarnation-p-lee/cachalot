package snapshot

import (
    "time"
)

// Snapshot indicates the timestamp information of host machine.
type Snapshot struct {
    Processes []Process
    Timestamp time.Time
}

// Process indicates the process related data.
type Process struct {
    CmdLine string
    CPU CPUStat
}

// CPUStat indicates the data for cpu stat.
type CPUStat struct {
    MCoreUsed, MCoreLimited float64
    UsageInPercentage float64
}

// CreateSnapshot will create one object with given timestamp.
func CreateSnapshot(timestamp time.Time) Snapshot {
    return Snapshot {
        Processes: []Process {},
        Timestamp: timestamp,
    }
}

// CreateProcess will create one object with given cmdLine.
func CreateProcess(cmdLine string) Process {
    return Process {
        CmdLine: cmdLine,
    }
}

// CreateCPUStat will create one object with cpu usage and limit, count in mCore.
func CreateCPUStat(mCoreUsed, mCoreLimited float64) CPUStat {
    usageInPercentage := (mCoreUsed / mCoreLimited) * 100.0

    return CPUStat {
        MCoreUsed: mCoreUsed,
        MCoreLimited: mCoreLimited,
        UsageInPercentage: usageInPercentage,
    }
}

// SetCPUStat will set the cpu usage.
func (process *Process) SetCPUStat(cpuStat CPUStat) {
    process.CPU = cpuStat
}

// AppendProcess will add given process to the process silic of snapshot.
func (snapshot *Snapshot) AppendProcess(process Process) {
    snapshot.Processes = append(snapshot.Processes, process)
}


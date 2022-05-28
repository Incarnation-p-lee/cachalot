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
    Cpu CpuStat
}

// CpuStat indicates the data for cpu stat.
type CpuStat struct {
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

// CreateCpuStat will create one object with cpu usage and limit, count in mCore.
func CreateCpuStat(mCoreUsed, mCoreLimited float64) CpuStat {
    usageInPercentage := (mCoreUsed / mCoreLimited) * 100.0

    return CpuStat {
        MCoreUsed: mCoreUsed,
        MCoreLimited: mCoreLimited,
        UsageInPercentage: usageInPercentage,
    }
}

// SetCpuStat will set the cpu usage.
func (process *Process) SetCpuStat(cpuStat CpuStat) {
    process.Cpu = cpuStat 
}

// AppendProcess will add given process to the process silic of snapshot.
func (snapshot *Snapshot) AppendProcess(process Process) {
    snapshot.Processes = append(snapshot.Processes, process)
}


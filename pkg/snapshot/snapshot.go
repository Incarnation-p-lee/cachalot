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
    PID int
    CmdLine string
    CPU CPUStat
}

// CPUStat indicates the data for cpu stat.
type CPUStat struct {
    JiffiesUsed, JiffiesInTotal int
    UsageInPercentage float64
}

// CreateSnapshot will create one object with given timestamp.
func CreateSnapshot(timestamp time.Time, processes []Process) Snapshot {
    return Snapshot {
        Timestamp: timestamp,
        Processes: processes,
    }
}

// CreateProcess will create one object with given cmdLine.
func CreateProcess(cmdLine string, pID int) Process {
    return Process {
        PID: pID,
        CmdLine: cmdLine,
    }
}

// CreateCPUStat will create one object with cpu usage and limit, count in mCore.
func CreateCPUStat(jiffiesUsed, jiffiesInTotal int) CPUStat {
    usageInPercentage := float64(jiffiesUsed) / float64(jiffiesInTotal) * 100.0

    return CPUStat {
        JiffiesUsed: jiffiesUsed,
        JiffiesInTotal: jiffiesInTotal,
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


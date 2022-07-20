package snapshot

import (
	"time"
)

// Snapshot indicates the timestamp information of host machine.
type Snapshot struct {
	Processes []Process
	Network   Network
	Timestamp time.Time
}

// Network indicates the network information of host machine.
type Network struct {
	INodeToTCP4 map[string]TCPConnection
	INodeToTCP6 map[string]TCPConnection
}

// TCPConnection indicates the connection information of tcp4.
type TCPConnection struct {
	INode         string
	State         string
	RemoteAddress string
	RemotePort    int
	UID           string
}

// Process indicates the process related data.
type Process struct {
	PID         int
	CmdLine     string
	CPUStat     CPUStat
	ThreadsStat ThreadsStat
	MemoryStat  MemoryStat
	NetworkStat ProcessNetworkStat
}

// ProcessNetworkStat indicates the process related network data.
type ProcessNetworkStat struct {
	TCP4Stat ProcessTCPStat
	TCP6Stat ProcessTCPStat
}

// ProcessTCPStat indicates the tcp related data.
type ProcessTCPStat struct {
	ConnectionCount        int
	ConnectionCountByState map[string]int
}

// CPUStat indicates the data for cpu stat.
type CPUStat struct {
	JiffiesUsed, JiffiesInTotal int
	UsageInPercentage           float64
}

// MemoryStat indicates the virtual memory stat
type MemoryStat struct {
	TotalMemoryInKB   int
	VMSizeInKB        int
	UsageInPercentage float64
	VMRSSInKB         int
	VMDataInKB        int
	VMStkInKB         int
	VMExeInKB         int
	VMLibInKB         int
}

// ThreadsStat indicates the data of thread stat.
type ThreadsStat struct {
	ThreadsCount int
}

const (
	// TCPUnknown indicates TCP connection state is unknown
	TCPUnknown     = "Unknown"
	// TCPEstablished indicates TCP connection is in established state
	TCPEstablished = "Established"
	// TCPSynSent indicates TCP connection is in syn sent state
	TCPSynSent     = "SynSent"
	// TCPSynRecv indicates TCP connection is in syn received state
	TCPSynRecv     = "SynRecv"
	// TCPFinWait1 indicates TCP connection is in final wait 1 state
	TCPFinWait1    = "FinWait1"
	// TCPFinWait2 indicates TCP connection is in final wait 2 state
	TCPFinWait2    = "FinWait2"
	// TCPTimeWait indicates TCP connection is in time wait state
	TCPTimeWait    = "TimeWait"
	// TCPClose indicates TCP connection is in close state
	TCPClose       = "Close"
	// TCPCloseWait indicates TCP connection is in close wait state
	TCPCloseWait   = "CloseWait"
	// TCPLastACK indicates TCP connection is in last ack state
	TCPLastACK     = "LastACK"
	// TCPListen indicates TCP connection is in listen state
	TCPListen      = "Listen"
	// TCPClosing indicates TCP connection is in closing state
	TCPClosing     = "Closing"
	// TCPNewSynRecv indicates tcp connection is in new sync received state
	TCPNewSynRecv  = "NewSynRecv"
)

var tcpStates = []string{
	TCPUnknown,
	TCPEstablished,
	TCPSynSent,
	TCPSynRecv,
	TCPFinWait1,
	TCPFinWait2,
	TCPTimeWait,
	TCPClose,
	TCPCloseWait,
	TCPLastACK,
	TCPListen,
	TCPClosing,
	TCPNewSynRecv,
}

// GetTCPStates will return the slice of all tcp states.
func GetTCPStates() []string {
	return tcpStates
}

// CreateCPUStat will create one object with cpu usage and limit, count in mCore.
func CreateCPUStat(jiffiesUsed, jiffiesInTotal int) CPUStat {
	usageInPercentage := float64(jiffiesUsed) / float64(jiffiesInTotal) * 100.0

	return CPUStat{
		JiffiesUsed:       jiffiesUsed,
		JiffiesInTotal:    jiffiesInTotal,
		UsageInPercentage: usageInPercentage,
	}
}

// SetCPUStat will set the cpu usage.
func (process *Process) SetCPUStat(cpuStat CPUStat) {
	process.CPUStat = cpuStat
}

// AppendProcess will add given process to the process silic of snapshot.
func (snapshot *Snapshot) AppendProcess(process Process) {
	snapshot.Processes = append(snapshot.Processes, process)
}

// AppendProcesses will add the given processes to the tail of snapshot.
func (snapshot *Snapshot) AppendProcesses(processes []Process) {
	if snapshot.Processes == nil {
		snapshot.Processes = []Process{}
	}

	snapshot.Processes = append(snapshot.Processes, processes...)
}

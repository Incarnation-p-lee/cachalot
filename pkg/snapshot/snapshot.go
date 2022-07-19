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
	TcpUnknown     = "Unknown"
	TcpEstablished = "Established"
	TcpSynSent     = "SynSent"
	TcpSynRecv     = "SynRecv"
	TcpFinWait1    = "FinWait1"
	TcpFinWait2    = "FinWait2"
	TcpTimeWait    = "TimeWait"
	TcpClose       = "Close"
	TcpCloseWait   = "CloseWait"
	TcpLastACK     = "LastACK"
	TcpListen      = "Listen"
	TcpClosing     = "Closing"
	TcpNewSynRecv  = "NewSynRecv"
)

var tcpStates = []string{
	TcpUnknown,
	TcpEstablished,
	TcpSynSent,
	TcpSynRecv,
	TcpFinWait1,
	TcpFinWait2,
	TcpTimeWait,
	TcpClose,
	TcpCloseWait,
	TcpLastACK,
	TcpListen,
	TcpClosing,
	TcpNewSynRecv,
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

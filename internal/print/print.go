package print

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/options"
	"log"
	"sort"
	"time"
)

const (
	jsonPrefix                  = ""
	jsonIndent                  = "  "
	printSeparatedLine          = "========================================================================="
	printSubTitleLine           = "-------------------------------------------------------------------------"
	tcpStateNoIndentMinimalSize = 11

	tcp4TypeName = "TCP4"
	tcp6TypeName = "TCP6"
)

func printSnapshotTitle(title string) {
	if len(title) == 0 {
		title = "Unknown title"
	}

	fmt.Println(printSeparatedLine)
	fmt.Printf("%s\n", title)
	fmt.Println(printSubTitleLine)
}

func printSnapshotTimestamp(timestamp time.Time) {
	fmt.Printf("Timestamp: %v\n", timestamp)
	fmt.Println(printSeparatedLine)
}

func printSnapshotNetworkTCPConnections(tcpStat snapshot.TCPStat, tcpType string) {
	fmt.Printf("%s-Connections\t%d\n", tcpType, tcpStat.ConnectionCount)
}

func printSnapshotNetworkTCPConnectionsByState(tcpStat snapshot.TCPStat, state, tcpType string) {
	printSnapshotConnetionState(tcpType, state)

	countByState := tcpStat.ConnectionCountByState

	if count, has := countByState[state]; has {
		fmt.Printf("\t%d", count)
	}

	fmt.Printf("\n")
}

func printSnapshotNetworkTCPConnectionsStates(tcpStat snapshot.TCPStat, tcpType string) {
	states := snapshot.GetTCPStates()

	for _, state := range states {
		printSnapshotNetworkTCPConnectionsByState(tcpStat, state, tcpType)
	}
}

func printSnapshotNetworkStat(networkStat snapshot.NetworkStat) {
	fmt.Println("Print snapshot network stat:")
	fmt.Println(printSubTitleLine)

	printSnapshotNetworkTCPConnections(networkStat.TCP4Stat, tcp4TypeName)
	printSnapshotNetworkTCPConnectionsStates(networkStat.TCP4Stat, tcp4TypeName)

	printSnapshotNetworkTCPConnections(networkStat.TCP6Stat, tcp6TypeName)
	printSnapshotNetworkTCPConnectionsStates(networkStat.TCP6Stat, tcp6TypeName)

	fmt.Println(printSeparatedLine)
}

func printSnapshotStat(spshot snapshot.Snapshot) {
	printSnapshotNetworkStat(spshot.Network.NetworkStat)
}

func printSnapshotProcessesPID(processes []snapshot.Process) {
	fmt.Printf("PID\t\t")

	for _, process := range processes {
		fmt.Printf("\t%v", process.PID)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesCPUUsage(processes []snapshot.Process) {
	fmt.Printf("CPUUsage\t")

	for _, process := range processes {
		fmt.Printf("\t%.1f%%", process.CPUStat.UsageInPercentage)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesThreads(processes []snapshot.Process) {
	fmt.Printf("Threads\t\t")

	for _, process := range processes {
		fmt.Printf("\t%d", process.ThreadsStat.ThreadsCount)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesMemoryUsage(processes []snapshot.Process) {
	fmt.Printf("MemoryUsage\t")

	for _, process := range processes {
		fmt.Printf("\t%.1f%%", process.MemoryStat.UsageInPercentage)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesMemoryVMSize(processes []snapshot.Process) {
	fmt.Printf("VmSize\t\t")

	for _, process := range processes {
		fmt.Printf("\t%dMB", process.MemoryStat.VMSizeInKB/1024)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesMemoryVMRSS(processes []snapshot.Process) {
	fmt.Printf("VmRss\t\t")

	for _, process := range processes {
		fmt.Printf("\t%dMB", process.MemoryStat.VMRSSInKB/1024)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesMemoryVMData(processes []snapshot.Process) {
	fmt.Printf("VmData\t\t")

	for _, process := range processes {
		fmt.Printf("\t%dMB", process.MemoryStat.VMDataInKB/1024)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesMemoryVMStk(processes []snapshot.Process) {
	fmt.Printf("VmStk\t\t")

	for _, process := range processes {
		fmt.Printf("\t%dMB", process.MemoryStat.VMStkInKB/1024)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesMemoryVMExe(processes []snapshot.Process) {
	fmt.Printf("VmExe\t\t")

	for _, process := range processes {
		fmt.Printf("\t%dMB", process.MemoryStat.VMExeInKB/1024)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesMemoryVMLib(processes []snapshot.Process) {
	fmt.Printf("VmLib\t\t")

	for _, process := range processes {
		fmt.Printf("\t%dMB", process.MemoryStat.VMLibInKB/1024)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesCmdLine(processes []snapshot.Process) {
	fmt.Printf("PID\t\tCmdLine\n")

	for _, process := range processes {
		fmt.Printf("%v\t\t%s\n", process.PID, process.CmdLine)
	}

	fmt.Println(printSeparatedLine)
}

func printSnapshotProcessesTCP4Connections(processes []snapshot.Process) {
	fmt.Printf("TCP4-Connections")

	for _, process := range processes {
		fmt.Printf("\t%d", process.NetworkStat.TCP4Stat.ConnectionCount)
	}

	fmt.Printf("\n")
}

func printSnapshotConnetionState(prefix string, state string) {
	fmt.Printf("%v-%v", prefix, state)

	if len(state) < tcpStateNoIndentMinimalSize {
		fmt.Printf("\t")
	}
}

func printSnapshotProcessesTCP4ConnectionsByState(processes []snapshot.Process, state string) {
	printSnapshotConnetionState("TCP4", state)

	for _, process := range processes {
		countByState := process.NetworkStat.TCP4Stat.ConnectionCountByState

		if count, has := countByState[state]; has {
			fmt.Printf("\t%d", count)
		}
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesTCP4ConnectionsStates(processes []snapshot.Process) {
	states := snapshot.GetTCPStates()

	for _, state := range states {
		printSnapshotProcessesTCP4ConnectionsByState(processes, state)
	}
}

func printSnapshotProcessesTCP6Connections(processes []snapshot.Process) {
	fmt.Printf("TCP6-Connections")

	for _, process := range processes {
		fmt.Printf("\t%d", process.NetworkStat.TCP6Stat.ConnectionCount)
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesTCP6ConnectionsByState(processes []snapshot.Process, state string) {
	printSnapshotConnetionState("TCP6", state)

	for _, process := range processes {
		countByState := process.NetworkStat.TCP6Stat.ConnectionCountByState

		if count, has := countByState[state]; has {
			fmt.Printf("\t%d", count)
		}
	}

	fmt.Printf("\n")
}

func printSnapshotProcessesTCP6ConnectionsStates(processes []snapshot.Process) {
	states := snapshot.GetTCPStates()

	for _, state := range states {
		printSnapshotProcessesTCP6ConnectionsByState(processes, state)
	}
}

func printSnapshotProcessesNetwork(processes []snapshot.Process) {
	printSnapshotProcessesTCP4Connections(processes)
	printSnapshotProcessesTCP4ConnectionsStates(processes)

	printSnapshotProcessesTCP6Connections(processes)
	printSnapshotProcessesTCP6ConnectionsStates(processes)
}

func printSnapshotProcessesCPU(processes []snapshot.Process) {
	printSnapshotProcessesPID(processes)
	printSnapshotProcessesThreads(processes)
	printSnapshotProcessesCPUUsage(processes)
}

func printSnapshotProcessesMemory(processes []snapshot.Process) {
	printSnapshotProcessesMemoryUsage(processes)
	printSnapshotProcessesMemoryVMSize(processes)
	printSnapshotProcessesMemoryVMRSS(processes)
	printSnapshotProcessesMemoryVMData(processes)
	printSnapshotProcessesMemoryVMStk(processes)
	printSnapshotProcessesMemoryVMExe(processes)
	printSnapshotProcessesMemoryVMLib(processes)
}

func printSnapshotProcessesStat(processes []snapshot.Process) {
	printSnapshotProcessesCmdLine(processes)
	printSnapshotProcessesCPU(processes)
	printSnapshotProcessesMemory(processes)
	printSnapshotProcessesNetwork(processes)
}

func printSnapshotFoot() {
	fmt.Printf("\n\n")
}

func printTextSnapshot(snapshot snapshot.Snapshot, title string) {
	printSnapshotTitle(title)

	printSnapshotTimestamp(snapshot.Timestamp)
	printSnapshotStat(snapshot)
	printSnapshotProcessesStat(snapshot.Processes)

	printSnapshotFoot()
}

func printJSONSnapshot(snapshot snapshot.Snapshot) {
	// It is not easy to get errors when serialize object to string. Thus, ignore the error.
	jsonBytes, _ := json.MarshalIndent(snapshot, jsonPrefix, jsonIndent)

	fmt.Printf("%s\n", string(jsonBytes))
}

func reconcileSnapshotTopCount(snapshot *snapshot.Snapshot, topCount int) {
	if len(snapshot.Processes) > topCount {
		snapshot.Processes = snapshot.Processes[:topCount]
	}
}

func reconcileSnapshotSortedBy(snapshot *snapshot.Snapshot, sortedBy string) {
	switch sortedBy {
	case options.SortedByMemory:
		sort.Slice(snapshot.Processes, func(a, b int) bool {
			memoryA, memoryB := snapshot.Processes[a].MemoryStat, snapshot.Processes[b].MemoryStat
			return memoryA.UsageInPercentage > memoryB.UsageInPercentage
		})
	case options.SortedByThreadsCount:
		sort.Slice(snapshot.Processes, func(a, b int) bool {
			threadA, threadB := snapshot.Processes[a].ThreadsStat, snapshot.Processes[b].ThreadsStat
			return threadA.ThreadsCount > threadB.ThreadsCount
		})
	case options.SortedByCPU:
		fallthrough
	default:
		sort.Slice(snapshot.Processes, func(a, b int) bool {
			cpuA, cpuB := snapshot.Processes[a].CPUStat, snapshot.Processes[b].CPUStat
			return cpuA.UsageInPercentage > cpuB.UsageInPercentage
		})
	}
}

func reconcileSnapshot(snapshot *snapshot.Snapshot, ops *options.Options) {
	reconcileSnapshotTopCount(snapshot, ops.GetTopCount())
	reconcileSnapshotSortedBy(snapshot, ops.GetSortedBy())
}

// Snapshot will print the data module of given snapshot.
func Snapshot(snapshot snapshot.Snapshot, title string, ops *options.Options) error {
	if ops == nil {
		return errors.New("found nil ops for snapshot print, will do nothing here")
	}

	reconcileSnapshot(&snapshot, ops)
	outputFormat := ops.GetOutputFormat()

	switch outputFormat {
	case options.TextOutput:
		printTextSnapshot(snapshot, title)
	case options.JSONOutput:
		printJSONSnapshot(snapshot)
	default:
		log.Printf("Unknown output format %v, fall back to %+v.\n",
			outputFormat, options.TextOutput)
		printTextSnapshot(snapshot, title)
	}

	return nil
}

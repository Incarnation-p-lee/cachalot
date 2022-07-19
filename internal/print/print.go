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
	jsonPrefix         = ""
	jsonIndent         = "  "
	printSeparatedLine = "========================================================================="
)

func printSnapshotTitle(title string) {
	if len(title) == 0 {
		title = "Unknown title"
	}

	fmt.Println(printSeparatedLine)
	fmt.Printf("%s\n", title)
}

func printSnapshotTimestamp(timestamp time.Time) {
	fmt.Printf("Timestamp: %v\n", timestamp)
	fmt.Println(printSeparatedLine)
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

func printSnapshotProcessesTCP4ConnectionsByState(processes []snapshot.Process, state string) {
	fmt.Printf("TCP4-%v", state)

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

func printSnapshotProcessesNetwork(processes []snapshot.Process) {
	printSnapshotProcessesTCP4Connections(processes)
	printSnapshotProcessesTCP4ConnectionsStates(processes)
}

func printSnapshotProcesses(processes []snapshot.Process) {
	printSnapshotProcessesCmdLine(processes)
	printSnapshotProcessesPID(processes)
	printSnapshotProcessesThreads(processes)
	printSnapshotProcessesCPUUsage(processes)
	printSnapshotProcessesMemoryUsage(processes)
	printSnapshotProcessesMemoryVMSize(processes)
	printSnapshotProcessesMemoryVMRSS(processes)
	printSnapshotProcessesMemoryVMData(processes)
	printSnapshotProcessesMemoryVMStk(processes)
	printSnapshotProcessesMemoryVMExe(processes)
	printSnapshotProcessesMemoryVMLib(processes)
	printSnapshotProcessesNetwork(processes)
}

func printSnapshotFoot() {
	fmt.Printf("\n\n")
}

func printTextSnapshot(snapshot snapshot.Snapshot, title string) {
	printSnapshotTitle(title)

	printSnapshotTimestamp(snapshot.Timestamp)
	printSnapshotProcesses(snapshot.Processes)

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

package print

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/options"
	"log"
	"time"
)

const (
	jsonPrefix         = ""
	jsonIndent         = "  "
	printSeparatedLine = "=========================================================="
)

func printSnapshotTitle(title string) {
	if len(title) == 0 {
		title = "Unknown title"
	}

	fmt.Println(printSeparatedLine)
	fmt.Printf("%s\n", title)
	fmt.Println(printSeparatedLine)
}

func printSnapshotTimestamp(timestamp time.Time) {
	fmt.Printf("Timestamp: %v\n", timestamp)
	fmt.Println(printSeparatedLine)
}

func printSnapshotProcess(process snapshot.Process) {
	fmt.Printf("%v\t", process.PID)

	fmt.Printf("%.3f%%\t\t", process.CPUStat.UsageInPercentage)
	fmt.Printf("%d\t", process.ThreadsStat.ThreadsCount)

	fmt.Printf("%.3f%%\t\t", process.MemoryStat.UsageInPercentage)
	fmt.Printf("%dMB\t", process.MemoryStat.VMSizeInKB/1024)
	fmt.Printf("%dMB\t", process.MemoryStat.VMRSSInKB/1024)
	fmt.Printf("%dMB\t", process.MemoryStat.VMDataInKB/1024)
	fmt.Printf("%dMB\t", process.MemoryStat.VMStkInKB/1024)
	fmt.Printf("%dMB\t", process.MemoryStat.VMExeInKB/1024)
	fmt.Printf("%dMB\t", process.MemoryStat.VMLibInKB/1024)

	fmt.Printf("%s\n", process.CmdLine)
}

func printSnapshotProcesses(processes []snapshot.Process) {
	fmt.Printf("Total procesess count: %d\n", len(processes))
	fmt.Println(printSeparatedLine)
	fmt.Printf("PID\tCPUUsage\tThreads\t" +
		"MemoryUsage\tVmSize\tVmRss\tVmData\tVmStk\tVmExe\tVmLib\t" +
		"CmdLine\n")
	fmt.Println(printSeparatedLine)

	for _, process := range processes {
		printSnapshotProcess(process)
	}
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

// Snapshot will print the data module of given snapshot.
func Snapshot(snapshot snapshot.Snapshot, title string, ops *options.Options) error {
	if ops == nil {
		return errors.New("found nil ops for snapshot print, will do nothing here")
	}

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

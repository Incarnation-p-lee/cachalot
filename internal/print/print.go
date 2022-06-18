package print

import (
    "fmt"
    "log"
    "time"
    "errors"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func printSnapshotTitle(title string) {
    if len(title) == 0 {
        title = "Unknown title"
    }

    fmt.Printf("==========================================================\n")
    fmt.Printf("%s\n", title)
    fmt.Printf("==========================================================\n")
}

func printSnapshotTimestamp(timestamp time.Time) {
    fmt.Printf("Timestamp: %v\n", timestamp)
    fmt.Printf("==========================================================\n")
}

func printSnapshotProcess(process snapshot.Process) {

    fmt.Printf("%v\t\t", process.PID)
    fmt.Printf("%.3f%%\t\t", process.CPUStat.UsageInPercentage)
    fmt.Printf("%s\n", process.CmdLine)
}

func printSnapshotProcesses(processes []snapshot.Process) {
    fmt.Printf("Total procesess count: %d\n", len(processes))
    fmt.Printf("==========================================================\n")
    fmt.Printf("PID\t\tCPUUsage\tCmdLine\n")
    fmt.Printf("----------------------------------------------------------\n")

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

// Snapshot will print the data module of given snapshot.
func Snapshot(snapshot snapshot.Snapshot, title string, ops *options.Options) error {
    if ops == nil {
        return errors.New("found nil ops for snapshot print, will do nothing here")
    }

    outputFormat := ops.GetOutputFormat()

    switch (outputFormat) {
        case options.TextOutput:
            printTextSnapshot(snapshot, title)
        default:
            log.Printf("Unknown output format %v, fall back to %+v.\n",
                outputFormat, options.TextOutput)
    }

    return nil
}


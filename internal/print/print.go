package print

import (
    "fmt"
    "time"
    "errors"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func printSnapshotTitle(title string) {
    if len(title) == 0 {
        title = "Unknown title"
    }

    fmt.Printf("==================== %s ====================\n", title)
}

func printSnapshotTimestamp(timestamp time.Time) {
    fmt.Printf("Sampling timestamp %v.\n", timestamp)
}

func printSnapshotProcess(process snapshot.Process) {
    fmt.Printf("------------------------------------\n")

    fmt.Printf("%s:\t\t\t%v\n", "PID", process.PID)
    fmt.Printf("%s:\t\t%v\n", "CmdLine", process.CmdLine)
    fmt.Printf("%s:\t%v\n", "UsageInPercentage", process.CPUStat.UsageInPercentage)

    fmt.Printf("\n")
}

func printSnapshotProcesses(processes []snapshot.Process) {
    fmt.Printf("There are %d processes in total.\n", len(processes))

    for _, process := range processes {
        printSnapshotProcess(process)
    }
}

func printSnapshotFoot() {
    fmt.Printf("\n\n")
}

// Snapshot will print the data module of given snapshot.
func Snapshot(snapshot snapshot.Snapshot, title string, ops *options.Options) error {
    if ops == nil {
        return errors.New("found nil ops for snapshot print, will do nothing here")
    }

    printSnapshotTitle(title)

    printSnapshotTimestamp(snapshot.Timestamp)
    printSnapshotProcesses(snapshot.Processes)

    printSnapshotFoot()

    return nil
}


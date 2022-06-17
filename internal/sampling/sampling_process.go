package sampling

import (
    "log"
    "sort"
    "path"
    "strconv"
    "path/filepath"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func getAllProcessID() []int {
    allPIDs, processPattern := []int {}, "/proc/[0-9]*"
    files, err := filepath.Glob(processPattern)

    if err != nil {
        log.Printf("Failed to open dir %s due to %+v\n", processPattern, err)
        return allPIDs
    }

    for _, file := range files {
        if pID, err := strconv.Atoi(path.Base(file)); err == nil {
            allPIDs = append(allPIDs, pID)
        }
    }

    return allPIDs
}

func sampleAllProcess(ops *options.Options) []snapshot.Process {
    allPIDs := getAllProcessID()

    pIDCount := len(allPIDs)
    pIDChan, processChan := make(chan int, pIDCount), make(chan snapshot.Process, pIDCount)

    for i := 0; i < pIDCount; i++ {
        go sampleOneProcessSnapshot(ops, pIDChan, processChan)
    }

    for _, pID := range allPIDs {
        pIDChan <- pID
    }

    processes := []snapshot.Process {}

    for i := 0; i < pIDCount; i++ {
        processes = append(processes, <- processChan)
    }

    sort.Slice(processes, func(a, b int) bool {
        return processes[a].CPUStat.UsageInPercentage > processes[b].CPUStat.UsageInPercentage
    })

    return processes
}

func sampleOneProcessSnapshot(ops *options.Options, pIDChan <- chan int,
    processChan chan <- snapshot.Process) {

    if ops == nil {
        processChan <- snapshot.Process {}
        return
    }

    pID := <- pIDChan
    cmdLine, cpuStat := sampleCmdLine(pID), sampleCPUStat(pID)

    processChan <- snapshot.CreateProcess(pID, cmdLine, cpuStat)
}


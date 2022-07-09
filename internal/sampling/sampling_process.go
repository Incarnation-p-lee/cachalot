package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/options"
	"log"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func getAllProcessIDs() []int {
	allPIDs, processPattern := []int{}, "/proc/[0-9]*"
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

func getProcessIDs(processStringIDs string) []int {
	stringIDs := strings.Split(processStringIDs, ",")
	processIDs := []int{}

	for _, v := range stringIDs {
		if pID, err := strconv.Atoi(v); err == nil {
			processIDs = append(processIDs, pID)
		}
	}

	return processIDs
}

func getSamplingProcessIDs(ops *options.Options) []int {
	if ops.IsAllProcessIDs() {
		return getAllProcessIDs()
	}

	return getProcessIDs(ops.GetProcessIDs())
}

func sampleProcesses(ops *options.Options) []snapshot.Process {
	allPIDs := getSamplingProcessIDs(ops)

	pIDCount := len(allPIDs)
	processChan := make(chan snapshot.Process, pIDCount)

	defer close(processChan)

	for _, pID := range allPIDs {
		go sampleOneProcessSnapshot(ops, pID, processChan)
	}

	processes := []snapshot.Process{}

	for i := 0; i < pIDCount; i++ {
		processes = append(processes, <-processChan)
	}

	return processes
}

func sampleOneProcessSnapshot(ops *options.Options, pID int,
	processChan chan<- snapshot.Process) {

	if ops == nil {
		processChan <- snapshot.Process{}
		return
	}

	cmdChan, cpuChan := make(chan string), make(chan snapshot.CPUStat)
	threadsChan, memoryChan := make(chan snapshot.ThreadsStat), make(chan snapshot.MemoryStat)

	defer close(cmdChan)
	defer close(cpuChan)
	defer close(threadsChan)
	defer close(memoryChan)

	go sampleCmdLine(pID, cmdChan)
	go sampleCPUStat(pID, cpuChan)
	go sampleThreadsStat(pID, threadsChan)
	go sampleMemoryStat(pID, memoryChan)

	processChan <- snapshot.Process{
		PID:         pID,
		CmdLine:     <-cmdChan,
		CPUStat:     <-cpuChan,
		ThreadsStat: <-threadsChan,
		MemoryStat:  <-memoryChan,
	}
}

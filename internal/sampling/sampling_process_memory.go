package sampling

import (
	"bufio"
	"fmt"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	invalidMemoryInKB = -1
	invalidMemoryName = "unknown memory type"
	totalMemoryFile   = "/proc/meminfo"
	totalMemoryPrefix = "MemTotal:"
)

var totalMemoryInKB = invalidMemoryInKB
var invalidMemoryStat = snapshot.MemoryStat{
	TotalMemoryInKB: invalidMemoryInKB,
	VMSizeInKB:      invalidMemoryInKB,
	VMRSSInKB:       invalidMemoryInKB,
	VMDataInKB:      invalidMemoryInKB,
	VMStkInKB:       invalidMemoryInKB,
	VMExeInKB:       invalidMemoryInKB,
	VMLibInKB:       invalidMemoryInKB,
}

func getMemoryInKB(memoryLine string) int {
	memoryInKB := getFirstIntValue(memoryLine)

	if memoryInKB == invalidSamplingIntValue {
		memoryInKB = invalidMemoryInKB
	}

	return memoryInKB
}

func getMemoryName(memoryLine string) string {
	memoryName := getFirstStringValue(memoryLine)

	if memoryName == invalidSamplingStringValue {
		memoryName = invalidMemoryName
	}

	return memoryName
}

func initTotalMemoryInKB() {
	file, err := os.Open(totalMemoryFile)

	if err != nil {
		log.Printf("Failed to open file %s due to %+v\n", totalMemoryFile, err)
		return
	}

	defer utils.CloseFile(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if line := scanner.Text(); strings.HasPrefix(line, totalMemoryPrefix) {
			totalMemoryInKB = getMemoryInKB(line)
			break
		}
	}
}

func sampleMemoryStat(pID int, memoryStatChan chan<- snapshot.MemoryStat) {
	filename := fmt.Sprintf("/proc/%d/status", pID)
	file, err := os.Open(filepath.Clean(filename))

	if err != nil {
		log.Printf("Failed to open file %s due to %+v\n", filename, err)
		memoryStatChan <- invalidMemoryStat
		return
	}

	defer utils.CloseFile(file)

	scanner, memoryStat := bufio.NewScanner(file), invalidMemoryStat

	for scanner.Scan() {
		line := scanner.Text()
		key := getMemoryName(line)

		switch key {
		case "VmSize":
			memoryStat.VMSizeInKB = getMemoryInKB(line)
		case "VmRSS":
			memoryStat.VMRSSInKB = getMemoryInKB(line)
		case "VmData":
			memoryStat.VMDataInKB = getMemoryInKB(line)
		case "VmStk":
			memoryStat.VMStkInKB = getMemoryInKB(line)
		case "VmExe":
			memoryStat.VMExeInKB = getMemoryInKB(line)
		case "VmLib":
			memoryStat.VMLibInKB = getMemoryInKB(line)
		}
	}

	memoryStat.TotalMemoryInKB = totalMemoryInKB
	memoryStat.UsageInPercentage =
		float64(memoryStat.VMSizeInKB*100) / float64(totalMemoryInKB)

	memoryStatChan <- memoryStat
}

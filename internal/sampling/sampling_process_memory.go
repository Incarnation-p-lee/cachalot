package sampling

import (
	"bufio"
	"fmt"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/utils"
	"log"
	"path/filepath"
	"os"
	"strings"
)

const (
	invalidMemoryInKB = -1
	invalidMemoryName = "unknown memory type"
	totalMemoryFile   = "/proc/meminfo"
	totalMemoryPrefix = "MemTotal:"
)

var totalMemoryInKB = invalidMemoryInKB
var invalidMemoryStat = snapshot.MemoryStat {
	TotalMemoryInKB: invalidMemoryInKB,
	VMSizeInKB: invalidMemoryInKB,
	VMRSSInKB: invalidMemoryInKB,
	VMDataInKB: invalidMemoryInKB,
	VMStkInKB: invalidMemoryInKB,
	VMExeInKB: invalidMemoryInKB,
	VMLibInKB: invalidMemoryInKB,
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

func sampleMemoryStat(pID int) snapshot.MemoryStat {
	filename := fmt.Sprintf("/proc/%d/status", pID)
	file, err := os.Open(filepath.Clean(filename))

	if err != nil {
		log.Printf("Failed to open file %s due to %+v\n", filename, err)
		return invalidMemoryStat
	}

	defer utils.CloseFile(file)

	scanner := bufio.NewScanner(file)
	vmSize, vmRSS, vmData := invalidMemoryInKB, invalidMemoryInKB, invalidMemoryInKB
	vmStk, vmExe, vmLib := invalidMemoryInKB, invalidMemoryInKB, invalidMemoryInKB

	for scanner.Scan() {
		line := scanner.Text()
		key := getMemoryName(line)

		switch key {
			case "VmSize":
				vmSize = getMemoryInKB(line)
			case "VmRSS":
				vmRSS = getMemoryInKB(line)
			case "VmData":
				vmData = getMemoryInKB(line)
			case "VmStk":
				vmStk = getMemoryInKB(line)
			case "VmExe":
				vmExe = getMemoryInKB(line)
			case "VmLib":
				vmLib = getMemoryInKB(line)
		}
	}

	return snapshot.MemoryStat {
		TotalMemoryInKB: totalMemoryInKB,
		VMSizeInKB: vmSize,
		UsageInPercentage: float64(vmSize * 100) / float64(totalMemoryInKB),
		VMRSSInKB: vmRSS,
		VMDataInKB: vmData,
		VMStkInKB: vmStk,
		VMExeInKB: vmExe,
		VMLibInKB: vmLib,
	}
}

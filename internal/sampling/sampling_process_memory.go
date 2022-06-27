package sampling

import (
	"bufio"
	"internal/utils"
	"log"
	"os"
	"strings"
)

const (
	invalidMemoryInKB = -1
	totalMemoryFile = "/proc/meminfo"
	totalMemoryPrefix = "MemTotal:"
)

var totalMemoryInKB int = invalidMemoryInKB

func getTotalMemoryInKB(totalMemoryLine string) int {
	memoryInKB := getFirstIntValue(totalMemoryLine)

	if memoryInKB == invalidSamplingIntValue {
		memoryInKB = invalidMemoryInKB
	}

	return memoryInKB
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
			totalMemoryInKB = getTotalMemoryInKB(line)
			break
		}
	}
}


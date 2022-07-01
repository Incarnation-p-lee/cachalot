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
	threadsCountPrefix  = "Threads"
	invalidThreadsCount = -1
)

func getThreadsCount(threadsLine string) int {
	threadsCount := getFirstIntValue(threadsLine)

	if threadsCount == invalidSamplingIntValue {
		return invalidThreadsCount
	}

	return threadsCount
}

func sampleThreadsCount(pID int) int {
	filename := fmt.Sprintf("/proc/%d/status", pID)
	file, err := os.Open(filepath.Clean(filename))

	if err != nil {
		log.Printf("Failed to open file %s due to %+v\n", filename, err)
		return invalidThreadsCount
	}

	defer utils.CloseFile(file)

	scanner, threadCount := bufio.NewScanner(file), invalidThreadsCount

	for scanner.Scan() {
		if line := scanner.Text(); strings.HasPrefix(line, threadsCountPrefix) {
			threadCount = getThreadsCount(line)
			break
		}
	}

	return threadCount
}

func sampleThreadsStat(pID int, threadsStatChan chan<- snapshot.ThreadsStat) {
	threadsStatChan<- snapshot.ThreadsStat{
		ThreadsCount: sampleThreadsCount(pID),
	}
}

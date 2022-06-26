package sampling

import (
	"bufio"
	"fmt"
	"internal/utils"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	threadsCountPrefix  = "Threads"
	invalidThreadsCount = -1
)

func getThreadsCount(threadsLine string) int {
	threads, threadsCount := strings.Split(threadsLine, ":"), invalidThreadsCount
	threads = threads[1:] // skip leading 'Threads'

	for _, v := range threads {
		v := strings.Trim(v, " \t")

		if len(v) > 0 {
			if count, err := strconv.Atoi(v); err != nil {
				log.Printf("Failed to convert integer from %s due to %+v\n", v, err)
			} else {
				threadsCount = count
				break
			}
		}
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

func sampleThreadStat(pID int) snapshot.ThreadsStat {
	return snapshot.ThreadsStat{
		ThreadsCount: sampleThreadsCount(pID),
	}
}

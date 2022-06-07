package sampling

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "io/ioutil"
    "path/filepath"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

const (
    totalStatFile = "/proc/stat"
    cpuNamePrefix = "cpu"

    invalidJiffies = -1
)

func getTotalCPUJiffies() int {
    file, err := os.Open(totalStatFile)

    if err != nil {
        log.Printf("Failed to open file %s due to %+v\n", totalStatFile, err)
        return invalidJiffies
    }

    defer file.Close()

    scanner, jiffies := bufio.NewScanner(file), 0

    for scanner.Scan() {
        if line := scanner.Text(); strings.HasPrefix(line, cpuNamePrefix) {
            jiffies += getOneCPUJiffies(line)
        }
    }

    return jiffies
}

func getOneCPUJiffies(cpuLine string) int {
    jiffies, total := strings.Split(cpuLine, " "), 0
    jiffies = jiffies[1:] // skip leading cpu[x]

    for _, v := range jiffies {
        if len(v) > 0 {
            if count, err := strconv.Atoi(v); err != nil {
                log.Printf("Failed to convert integer from %s due to %+v\n", v, err)
            } else {
                total += count
            }
        }
    }

    return total
}

func getProcessCPUJiffies(pId int) int {
    file := fmt.Sprintf("/proc/%d/stat", pId)
    _, err := ioutil.ReadFile(filepath.Clean(file))

    if err != nil {
        log.Printf("Failed to open file %s due to %+v\n", file, err)
        return invalidJiffies
    }

    return invalidJiffies
}

func sampleCPU(pId int) snapshot.CPUStat {
    totalJiffies := getTotalCPUJiffies()
    processJiffies := getProcessCPUJiffies(pId)

    return snapshot.CreateCPUStat(processJiffies, totalJiffies)
}


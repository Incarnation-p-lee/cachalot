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

    defaultJiffies = 0

    userJiffiesIndex = 14
    kernelJiffiesIndex = 15
    childrenUserJiffiesIndex = 16
    childrenKernelJiffiesIndex = 17

    jiffiesMaxSize = childrenKernelJiffiesIndex + 1
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

func getJiffiesOrDefault(stats []string, index int) int {
    val := stats[index]
    jiffies, err := strconv.Atoi(val)

    if err != nil {
        log.Printf("Failed to convert integer from stats[%d](%s) due to %+v\n", index, val, err)

        return defaultJiffies
    }

    return jiffies
}

func getProcessCPUJiffies(pID int) int {
    file := fmt.Sprintf("/proc/%d/stat", pID)
    content, err := ioutil.ReadFile(filepath.Clean(file))

    if err != nil {
        log.Printf("Failed to open file %s due to %+v\n", file, err)
        return invalidJiffies
    }

    stats := strings.Split(string(content), " ")

    if len(stats) < jiffiesMaxSize {
        log.Printf("Stats slice size should be greater than %d.\n", jiffiesMaxSize)
        return invalidJiffies
    }

    userJiffies := getJiffiesOrDefault(stats, userJiffiesIndex)
    kernelJiffies := getJiffiesOrDefault(stats, kernelJiffiesIndex)
    childrenUserJiffies := getJiffiesOrDefault(stats, childrenUserJiffiesIndex)
    childrenKernelJiffies := getJiffiesOrDefault(stats, childrenKernelJiffiesIndex)

    return userJiffies + kernelJiffies + childrenUserJiffies + childrenKernelJiffies
}

func sampleCPU(pID int) snapshot.CPUStat {
    totalJiffies := getTotalCPUJiffies()
    processJiffies := getProcessCPUJiffies(pID)

    return snapshot.CreateCPUStat(processJiffies, totalJiffies)
}


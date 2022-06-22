package sampling

import (
    "os"
    "log"
    "fmt"
    "time"
    "bufio"
    "strings"
    "strconv"
    "io/ioutil"
    "path/filepath"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

const (
    totalStatFile = "/proc/stat"
    cpuNamePrefix = "cpu "

    invalidJiffies = -1
    defaultJiffies = 0

    samplingDuration = 5 // Count in seconds

    userJiffiesIndex = 14 - 1
    kernelJiffiesIndex = 15 - 1
    childrenUserJiffiesIndex = 16 - 1
    childrenKernelJiffiesIndex = 17 - 1

    jiffiesMaxSize = childrenKernelJiffiesIndex + 1
)

func sampleTotalCPUJiffies(totalChan chan <- int) {
    file, err := os.Open(totalStatFile)

    if err != nil {
        log.Printf("Failed to open file %s due to %+v\n", totalStatFile, err)
        totalChan <- invalidJiffies
        return
    }

    defer func() {
        if err := file.Close(); err != nil {
            log.Printf("Failed to close file %+v due to %+v\n", file, err)
        }
    } ()

    scanner, jiffies := bufio.NewScanner(file), 0

    for scanner.Scan() {
        if line := scanner.Text(); strings.HasPrefix(line, cpuNamePrefix) {
            jiffies = getTotalCPUJiffies(line)
            break
        }
    }

    totalChan <- jiffies
}

func getTotalCPUJiffies(cpuLine string) int {
    jiffies, total := strings.Split(cpuLine, " "), 0
    jiffies = jiffies[1:] // skip leading 'cpu '

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

func sampleProcessCPUJiffies(processChan chan int) {
    pID := <- processChan
    file := fmt.Sprintf("/proc/%d/stat", pID)
    content, err := ioutil.ReadFile(filepath.Clean(file))

    if err != nil {
        log.Printf("Failed to open file %s due to %+v\n", file, err)
        processChan <- invalidJiffies
        return
    }

    stats := strings.Split(string(content), " ")

    if len(stats) < jiffiesMaxSize {
        log.Printf("Stats slice size should be greater than %d.\n", jiffiesMaxSize)
        processChan <- invalidJiffies
        return
    }

    userJiffies := getJiffiesOrDefault(stats, userJiffiesIndex)
    kernelJiffies := getJiffiesOrDefault(stats, kernelJiffiesIndex)
    childrenUserJiffies := getJiffiesOrDefault(stats, childrenUserJiffiesIndex)
    childrenKernelJiffies := getJiffiesOrDefault(stats, childrenKernelJiffiesIndex)

    allJiffies := userJiffies + kernelJiffies + childrenUserJiffies + childrenKernelJiffies
    processChan <- allJiffies
}

func sampleCPUJiffies(pID int) (totalJiffies, processJiffies int) {
    totalChan, processChan := make(chan int), make(chan int)

    defer close(totalChan)
    defer close(processChan)

    go sampleTotalCPUJiffies(totalChan)
    go sampleProcessCPUJiffies(processChan)

    processChan <- pID

    totalJiffies, processJiffies = <- totalChan, <- processChan

    return totalJiffies, processJiffies
}

func sampleCPUStat(pID int) snapshot.CPUStat {
    totalJiffiesBefore, processJiffiesBefore := sampleCPUJiffies(pID)

    time.Sleep(time.Duration(samplingDuration) * time.Second)

    totalJiffiesAfter, processJiffiesAfter := sampleCPUJiffies(pID)

    totalJiffies := totalJiffiesAfter - totalJiffiesBefore
    processJiffies := processJiffiesAfter - processJiffiesBefore

    return snapshot.CreateCPUStat(processJiffies, totalJiffies)
}


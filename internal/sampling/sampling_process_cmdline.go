package sampling

import (
    "log"
    "fmt"
    "io/ioutil"
    "path/filepath"
)

const (
    unknownCmdLine = "unknown cmd line"
)

func sampleCmdLine(pID int) string {
    file := fmt.Sprintf("/proc/%d/cmdline", pID)
    content, err := ioutil.ReadFile(filepath.Clean(file))

    if err != nil {
        log.Printf("Failed to open file %s due to %+v\n", file, err)
        return unknownCmdLine
    }

    return string(content)
}


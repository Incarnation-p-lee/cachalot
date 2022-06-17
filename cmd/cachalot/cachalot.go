package main

import (
    "fmt"
    "time"
    "internal/print"
    "internal/options"
    "internal/cmdflags"
    "internal/sampling"
)

func main() {
    ops := options.CreateOptions()

    cmdflags.ParseOptions(ops)

    sampleAndPrint(ops)
}

func sampleAndPrint(ops *options.Options) {
    count := ops.GetSamplingCount()
    interval := ops.GetSamplingInterval()

    for i := 0; i < count; i++ {
        snapshot := sampling.Sample(ops)
        title := fmt.Sprintf("Print snapshot with count %d.", i)

        print.Snapshot(snapshot, title, ops)

        time.Sleep(time.Duration(interval) * time.Second)
    }
}


package main

import (
    "time"
    "internal/options"
    "internal/cmdflags"
    "internal/sampling"
)

func main() {
    ops := options.CreateOptions()

    cmdflags.ParseOptions(ops)

    sample(ops)
}

func sample(ops *options.Options) {
    count := ops.GetSamplingCount()
    interval := ops.GetSamplingInterval()

    for i := 0; i < count; i++ {
        sampling.Sample(ops)

        time.Sleep(time.Duration(interval) * time.Second)
    }
}


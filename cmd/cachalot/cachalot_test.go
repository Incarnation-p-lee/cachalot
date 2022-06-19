package main

import (
    "time"
    "testing"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func getMinimalDuration(ops *options.Options) time.Duration {
    count := ops.GetSamplingCount()
    interval := ops.GetSamplingInterval()

    return time.Duration(interval * count) * time.Second
}

func TestSampleAndPrint(t *testing.T) {
    ops := options.CreateOptions()

    ops.AppendOption(options.Option {
        Key: options.SamplingCount,
        Val: "1",
    })
    ops.AppendOption(options.Option {
        Key: options.SamplingInterval,
        Val: "0",
    })

    start := time.Now()

    sampleAndPrint(ops)

    end := time.Now()

    assert.IsTrue(t, end.Sub(start) > getMinimalDuration(ops),
        "sample and print cost should be greater than minimal duration.")
}


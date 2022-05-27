package main

import (
    "internal/options"
    "internal/cmdflags"
)

func main() {
    ops := options.CreateOptions()

    cmdflags.ParseOptions(ops)
}


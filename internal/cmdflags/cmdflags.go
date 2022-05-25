package cmdflags

import (
    "flag"
    "internal/options"
)

type cmdflag struct {
    flagName, defaultValue string
    description string
}

var supportedCmdFlags = []cmdflag {
    cmdflag {
        flagName: options.SamplingCount,
        defaultValue: "10",
        description: "the total count of sampling",
    },
    cmdflag {
        flagName: options.SamplingInterval,
        defaultValue: "10",
        description: "the interval for each sampling, count in seconds",
    },
}

// ParseOptions will parse the flags from command line to options.
func ParseOptions(ops *options.Options) {
    if ops == nil {
        return
    }

    values := make([]string, len(supportedCmdFlags))

    for i, v := range supportedCmdFlags {
        flag.StringVar(&values[i], v.flagName, v.defaultValue, v.description)
    }

    flag.Parse()

    for i, v := range supportedCmdFlags {
        key, val := v.flagName, values[i]
        ops.AppendOption(options.CreateEnabledOption(key, val))
    }
}


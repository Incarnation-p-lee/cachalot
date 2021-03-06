package cmdflags

import (
	"flag"
	"fmt"
	"internal/options"
)

type cmdflag struct {
	flagName, defaultValue string
	description            string
}

var supportedCmdFlags = []cmdflag{
	cmdflag{
		flagName:     options.SamplingCount,
		defaultValue: options.GetNameDefaultValue(options.SamplingCount),
		description:  "the total count of sampling",
	},
	cmdflag{
		flagName:     options.SamplingInterval,
		defaultValue: options.GetNameDefaultValue(options.SamplingInterval),
		description:  "the interval for each sampling, count in seconds",
	},
	cmdflag{
		flagName:     options.OutputFormat,
		defaultValue: options.GetNameDefaultValue(options.OutputFormat),
		description: fmt.Sprintf("the output layout for print, supported formats are %+v",
			options.GetSupportedOutputFormatSlice()),
	},
	cmdflag{
		flagName:     options.ProcessIDs,
		defaultValue: options.GetNameDefaultValue(options.ProcessIDs),
		description:  "the comma separated pids for snapshot, -1 indicates all processes",
	},
	cmdflag{
		flagName:     options.TopCount,
		defaultValue: options.GetNameDefaultValue(options.TopCount),
		description:  "the top count of process to be printed",
	},
	cmdflag{
		flagName:     options.SortedBy,
		defaultValue: options.GetNameDefaultValue(options.SortedBy),
		description: fmt.Sprintf("the metrics to be sorted when print, supported metrics are %+v",
			options.GetSupportedSortedBySlice()),
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

package options

import (
	"errors"
	"fmt"
	"strconv"
)

// Option compose of key-value pair for one option, it also has one flag for enabled or not.
type Option struct {
	Key, Val string
	Enabled  bool
}

// Options is the collection of Option.
type Options struct {
	allOptions []Option
}

const (
	unknownDefaultValue = "unknown default value"

	// SamplingCount indicates how many count will be sampled.
	SamplingCount = "sampling-count"
	// SamplingInterval indicates the interval for each sampling, count in seconds.
	SamplingInterval = "sampling-interval"
	// OutputFormat indicates the layout when print.
	OutputFormat = "out"
	// ProcessIDs indicates the id of process.
	ProcessIDs = "pids"
	// TopCount indicates the limit of print times.
	TopCount = "top-count"
	// SortedBy indicates how to sort the print.
	SortedBy = "sorted-by"

	// AllProcessIDs indicates all process in a system.
	AllProcessIDs = "-1"
	// JSONOutput will be printed as json.
	JSONOutput = "json"
	// TextOutput will be printed as raw text.
	TextOutput = "text"
	// DefaultTopCount indicates the default print times.
	DefaultTopCount = "7"
	// SortedByCPU will sort the print by CPU.
	SortedByCPU = "cpu"
	// SortedByMemory will sort the print by memory.
	SortedByMemory = "memory"
)

var namesToDefaultValues = map[string]string{
	SamplingCount:    "10",
	SamplingInterval: "10",
	OutputFormat:     TextOutput,
	ProcessIDs:       AllProcessIDs,
	TopCount:         DefaultTopCount,
	SortedBy:         SortedByCPU,
}

// GetNameDefaultValue will return the default value for option name, or unknown.
func GetNameDefaultValue(name string) string {
	if value, has := namesToDefaultValues[name]; has {
		return value
	}

	return unknownDefaultValue
}

// CreateEnabledOption will create option enabled with given key.
func CreateEnabledOption(key, val string) Option {
	return Option{
		Key:     key,
		Val:     val,
		Enabled: true,
	}
}

// CreateOptions will create the object of Options and return the pointer.
func CreateOptions() *Options {
	return &Options{
		allOptions: []Option{},
	}
}

// AppendOption will append the Option to the Options.
func (ops *Options) AppendOption(option Option) {
	ops.allOptions = append(ops.allOptions, option)
}

func (ops *Options) getOptionVal(key string) string {
	val := ""

	for _, op := range ops.allOptions {
		if op.Key == key {
			val = op.Val
			break
		}
	}

	return val
}

func (ops *Options) getIntOption(key string) int {
	val := ops.getOptionVal(key)

	if intVal, err := strconv.Atoi(val); err == nil {
		return intVal
	}

	return 0
}

func (ops *Options) getStringOption(key string) string {
	return ops.getOptionVal(key)
}

// GetSamplingCount will return the sampling count, or zero if not found or any error.
func (ops *Options) GetSamplingCount() int {
	return ops.getIntOption(SamplingCount)
}

// GetSamplingInterval will return the sampling interval, or zero if not found or any error.
func (ops *Options) GetSamplingInterval() int {
	return ops.getIntOption(SamplingInterval)
}

// GetOutputFormat will return the output format, for example, json or text.
func (ops *Options) GetOutputFormat() string {
	return ops.getStringOption(OutputFormat)
}

// GetSupportedOutputFormatSlice will return all the output format supported.
func GetSupportedOutputFormatSlice() []string {
	return []string {
		TextOutput,
		JSONOutput,
	}
}

// GetProcessIDs will return the process ids.
func (ops *Options) GetProcessIDs() string {
	return ops.getStringOption(ProcessIDs)
}

// IsAllProcessIDs will return true if all proccess ids, or will return false.
func (ops *Options) IsAllProcessIDs() bool {
	return ops.getStringOption(ProcessIDs) == AllProcessIDs
}

// GetTopCount will return the top count of process to be print.
func (ops *Options) GetTopCount() int {
	return ops.getIntOption(TopCount)
}

// GetSortedBy will return the metrics sorted by.
func (ops *Options) GetSortedBy() string {
	return ops.getStringOption(SortedBy)
}

// GetSupportedSortedBySlice will return the metrics slice supported by sorted.
func GetSupportedSortedBySlice() []string {
	return []string {
		SortedByCPU,
		SortedByMemory,
	}
}

// GetOption will return the option by index, out of range will return error.
func (ops *Options) GetOption(index int) (Option, error) {
	limit := len(ops.allOptions)

	if index >= limit {
		msg := fmt.Sprintf("Requred index %d is out of limit %d.", index, limit)
		return Option{}, errors.New(msg)
	}

	return ops.allOptions[index], nil
}

// OptionsCount will return the total count of all Options.
func (ops *Options) OptionsCount() int {
	return len(ops.allOptions)
}

// IsSamplingCount indicates if the option is sampling count.
func (op *Option) IsSamplingCount() bool {
	return op.Key == SamplingCount
}

// IsSamplingInterval indicates if the option is sampling interval.
func (op *Option) IsSamplingInterval() bool {
	return op.Key == SamplingInterval
}

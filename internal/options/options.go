package options

import (
    "errors"
    "fmt"
    "strconv"
)

// Option compose of key-value pair for one option, it also has one flag for enabled or not.
type Option struct {
    Key, Val string
    Enabled bool
}

// Options is the collection of Option.
type Options struct {
    allOptions []Option
}

const (
    // SamplingCount indicates how many count will be sampled.
    SamplingCount = "sampling-count"
    // SamplingInterval indicates the interval for each sampling, count in seconds.
    SamplingInterval = "sampling-interval"
)

// CreateEnabledOption will create option enabled with given key.
func CreateEnabledOption(key, val string) Option {
    return Option {
        Key: key,
        Val: val,
        Enabled: true,
    }
}

// CreateOptions will create the object of Options and return the pointer.
func CreateOptions() *Options {
    return &Options {
        allOptions: []Option {},
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

    if intVal, err := strconv.Atoi(val); err != nil {
        return 0
    } else {
        return intVal
    }
}

// GetSamplingCount will return the sampling count, or zero if not found or any error.
func (ops *Options) GetSamplingCount() int {
    return ops.getIntOption(SamplingCount)
}

// GetSamplingInterval will return the sampling interval, or zero if not found or any error.
func (ops *Options) GetSamplingInterval() int {
    return ops.getIntOption(SamplingInterval)
}

// GetOption will return the option by index, out of range will return error.
func (ops *Options) GetOption(index int) (Option, error) {
    limit := len(ops.allOptions)

    if index >= limit {
        msg := fmt.Sprintf("Requred index %d is out of limit %d.", index, limit)
        return Option {}, errors.New(msg)
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


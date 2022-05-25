package options

import (
    "errors"
    "fmt"
)

type Option struct {
    Key, Val string
    Enabled bool
}

type Options struct {
    allOptions []Option
}

const (
    SamplingCount = "sampling-count"
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

func (op *Option) IsSamplingCount() bool {
    return op.Key == SamplingCount
}

func (op *Option) IsSamplingInterval() bool {
    return op.Key == SamplingInterval
}


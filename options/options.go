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

// CreateOptions will create the object of Options and return the pointer.
func CreateOptions() *Options {
    return &Options {
        allOptions: []Option {},
    }
}

// AppendOption will append the Option to the Options.
func (this *Options) AppendOption(option Option) {
    this.allOptions = append(this.allOptions, option)
}

// GetOption will return the option by index, out of range will return error.
func (this *Options) GetOption(index int) (Option, error) {
    limit := len(this.allOptions)

    if index >= limit {
        msg := fmt.Sprintf("Requred index %d is out of limit %d.", index, limit)
        return Option {}, errors.New(msg)
    }

    return this.allOptions[index], nil
}

// OptionsCount will return the total count of all Options.
func (this *Options) OptionsCount() int {
    return len(this.allOptions)
}


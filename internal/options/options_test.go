package options

import (
    "testing"
    "pkg/assert"
)

func TestCreateOptions(t *testing.T) {
    options := CreateOptions()

    assert.IsNotNil(t, options, "options cannot be null.")
    assert.IsNotNil(t, options.allOptions, "allOptions of options cannot be null.")
}

func TestAppendOption(t *testing.T) {
    options := CreateOptions()

    options.AppendOption(Option {})

    assert.IsEqual(t, 1, options.OptionsCount(), "optionsCount should be 1.")
}

func TestGetOption(t *testing.T) {
    options := CreateOptions()

    options.AppendOption(Option {
        Enabled: true,
    })

    options.AppendOption(Option {
        Enabled: false,
    })

    option, err := options.GetOption(0)

    assert.IsNil(t, err, "getOption should have no error.")
    assert.IsTrue(t, option.Enabled, "option should be enabled.")

    _, err = options.GetOption(2)

    assert.IsNotNil(t, err, "getOption should have error.")
}

func TestOptionsCount(t *testing.T) {
    options := CreateOptions()

    options.AppendOption(Option {})

    assert.IsEqual(t, 1, options.OptionsCount(), "optionsCount should be 1.")

    options.AppendOption(Option {})

    assert.IsEqual(t, 2, options.OptionsCount(), "optionsCount should be 2.")
}


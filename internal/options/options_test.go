package options

import (
    "testing"
    "pkg/assert"
)

func TestCreateEnabledOption(t *testing.T) {
    option := CreateEnabledOption("", "")

    assert.IsNotNil(t, option, "option should not be null.")
    assert.IsTrue(t, option.Enabled, "option should be enabled.")
}

func TestCreateOptions(t *testing.T) {
    options := CreateOptions()

    assert.IsNotNil(t, options, "options should not be null.")
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

func TestIsSamplingCount(t *testing.T) {
    countOption := Option { Key: SamplingCount, }
    intervalOption := Option { Key: SamplingInterval, }

    assert.IsTrue(t, countOption.IsSamplingCount(), "option should be sampling count.");
    assert.IsFalse(t, intervalOption.IsSamplingCount(), "option should not be sampling count.");
}

func TestIsSamplingInterval(t *testing.T) {
    countOption := Option { Key: SamplingCount, }
    intervalOption := Option { Key: SamplingInterval, }

    assert.IsTrue(t, intervalOption.IsSamplingInterval(), "option should be sampling interval.");
    assert.IsFalse(t, countOption.IsSamplingInterval(), "option should not be sampling interval.");
}


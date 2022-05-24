package options

import (
    "testing"
)

func TestCreateOptions(t *testing.T) {
    options := CreateOptions()

    if options == nil || options.allOptions == nil {
        t.Logf("Options cann be null.")
        t.Fail()
    }
}

func TestAppendOption(t *testing.T) {
    options := CreateOptions()

    options.AppendOption(Option {})

    if options.OptionsCount() != 1 {
        t.Logf("OptionsCount should be 1.")
        t.Fail()
    }
}

func TestGetOption(t *testing.T) {
    options := CreateOptions()

    options.AppendOption(Option {
        Enabled: true,
    })
    options.AppendOption(Option {
        Enabled: false,
    })

    if option, err := options.GetOption(0); err != nil {
        t.Logf("GetOption by 1 should has no error.")
        t.Fail()
    } else if option.Enabled == false {
        t.Logf("GetOption by 1 should be enabled.")
        t.Fail()
    }

    if _, err := options.GetOption(2); err == nil {
        t.Logf("GetOption by 2 should has error.")
        t.Fail()
    }
}

func TestOptionsCount(t *testing.T) {
    options := CreateOptions()

    options.AppendOption(Option {})

    if options.OptionsCount() != 1 {
        t.Logf("OptionCount should be 1.")
        t.Fail()
    }

    options.AppendOption(Option {})

    if options.OptionsCount() != 2 {
        t.Logf("OptionCount should be 1.")
        t.Fail()
    }
}


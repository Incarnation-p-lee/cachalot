package cmdflags

import (
    "testing"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestParseOptions(t *testing.T) {
    ops := options.CreateOptions()

    ParseOptions(ops)

    assert.IsEqual(t, 4, ops.OptionsCount(), "optionsCount should be 4.")

    option1, _ := ops.GetOption(0)
    option2, _ := ops.GetOption(1)

    assert.IsTrue(t, option1.IsSamplingCount(), "option is sampling count.")
    assert.IsTrue(t, option2.IsSamplingInterval(), "option is sampling interval");
}


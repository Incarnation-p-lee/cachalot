package options

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"strconv"
	"testing"
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

	options.AppendOption(Option{})

	assert.IsEqual(t, 1, options.OptionsCount(), "optionsCount should be 1.")
}

func TestGetOption(t *testing.T) {
	options := CreateOptions()

	options.AppendOption(Option{
		Enabled: true,
	})

	options.AppendOption(Option{
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

	options.AppendOption(Option{})

	assert.IsEqual(t, 1, options.OptionsCount(), "optionsCount should be 1.")

	options.AppendOption(Option{})

	assert.IsEqual(t, 2, options.OptionsCount(), "optionsCount should be 2.")
}

func TestIsSamplingCount(t *testing.T) {
	countOption := Option{Key: SamplingCount}
	intervalOption := Option{Key: SamplingInterval}

	assert.IsTrue(t, countOption.IsSamplingCount(), "option should be sampling count.")
	assert.IsFalse(t, intervalOption.IsSamplingCount(), "option should not be sampling count.")
}

func TestIsSamplingInterval(t *testing.T) {
	countOption := Option{Key: SamplingCount}
	intervalOption := Option{Key: SamplingInterval}

	assert.IsTrue(t, intervalOption.IsSamplingInterval(), "option should be sampling interval.")
	assert.IsFalse(t, countOption.IsSamplingInterval(), "option should not be sampling interval.")
}

func TestGetSamplingCountNormal(t *testing.T) {
	options := CreateOptions()

	options.AppendOption(Option{
		Key: SamplingCount,
		Val: "123",
	})

	assert.IsEqual(t, 123, options.GetSamplingCount(), "should have same sampling count")
}

func TestGetSamplingCountInvalid(t *testing.T) {
	options := CreateOptions()

	options.AppendOption(Option{
		Key: SamplingCount,
		Val: "invalid-number",
	})

	assert.IsEqual(t, 0, options.GetSamplingCount(), "should have same sampling count")
}

func TestGetSamplingIntervalNormal(t *testing.T) {
	options := CreateOptions()

	options.AppendOption(Option{
		Key: SamplingInterval,
		Val: "123",
	})

	assert.IsEqual(t, 123, options.GetSamplingInterval(), "should have same sampling count")
}

func TestGetSamplingIntervalInvalid(t *testing.T) {
	options := CreateOptions()

	options.AppendOption(Option{
		Key: SamplingInterval,
		Val: "invalid-number",
	})

	assert.IsEqual(t, 0, options.GetSamplingInterval(), "should have same sampling count")
}

func TestGetNameDefaultValue(t *testing.T) {
	assert.IsEqual(t, "10", GetNameDefaultValue(SamplingCount),
		"sampling count default value should be 10")

	assert.IsEqual(t, "unknown default value", GetNameDefaultValue("unknown"),
		"should have unknown default value")
}

func TestGetOutputFormat(t *testing.T) {
	ops := CreateOptions()

	assert.IsEqual(t, "", ops.GetOutputFormat(), "empty options should have empty output")

	ops.AppendOption(Option{
		Key: OutputFormat,
		Val: TextOutput,
	})

	assert.IsEqual(t, "text", ops.GetOutputFormat(), "options should have text output")
}

func TestGetProcessIds(t *testing.T) {
	ops, testPIDs := CreateOptions(), "1,2,3"

	ops.AppendOption(Option{
		Key: ProcessIDs,
		Val: testPIDs,
	})

	assert.IsEqual(t, testPIDs, ops.GetProcessIDs(), "options should have same process ids")
}

func TestIsAllProcessIDsTrue(t *testing.T) {
	ops := CreateOptions()

	ops.AppendOption(Option{
		Key: ProcessIDs,
		Val: AllProcessIDs,
	})

	assert.IsTrue(t, ops.IsAllProcessIDs(), "options should have all process ids")
}

func TestIsAllProcessIDsFalse(t *testing.T) {
	ops := CreateOptions()

	ops.AppendOption(Option{
		Key: ProcessIDs,
	})

	assert.IsFalse(t, ops.IsAllProcessIDs(), "options should not have all process ids")
}

func TestGetTopCount(t *testing.T) {
	ops, testCount := CreateOptions(), 2

	ops.AppendOption(Option{
		Key: TopCount,
		Val: strconv.Itoa(testCount),
	})

	assert.IsEqual(t, testCount, ops.GetTopCount(), "options should have same top count")
}

func TestGetSortedBy(t *testing.T) {
	ops := CreateOptions()

	ops.AppendOption(Option{
		Key: SortedBy,
		Val: SortedByMemory,
	})

	assert.IsEqual(t, "memory", ops.GetSortedBy(), "options should have same sorted by")
}

func TestGetSupportedSortedBySlice(t *testing.T) {
	supported := GetSupportedSortedBySlice()

	assert.IsEqual(t, 2, len(supported), "supported sorted by should be 2 in length")
	assert.IsEqual(t, "cpu", supported[0], "first supported sorted by should be cpu")
	assert.IsEqual(t, "memory", supported[1], "second supported sorted by should be memory")
}

func TestGetSupportedOutputFormatSlice(t *testing.T) {
	supported := GetSupportedOutputFormatSlice()

	assert.IsEqual(t, 2, len(supported), "supported output format should be 2 in length")
	assert.IsEqual(t, "text", supported[0], "first supported output format should be text")
	assert.IsEqual(t, "json", supported[1], "second supported output format should be json")
}

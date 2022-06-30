package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
)

func TestGetFirstIntValue(t *testing.T) {
	assert.IsEqual(t, invalidSamplingIntValue, getFirstIntValue(""),
		"empty line should be invalid sampling int value")
	assert.IsEqual(t, invalidSamplingIntValue, getFirstIntValue("abc"),
		"only chars line should be invalid sampling int value")
	assert.IsEqual(t, invalidSamplingIntValue, getFirstIntValue("a"),
		"only one space line should be invalid sampling int value")
	assert.IsEqual(t, invalidSamplingIntValue, getFirstIntValue("abc:def"),
		"no number line should be invalid sampling int value")

	assert.IsEqual(t, 123, getFirstIntValue("abc:123"),
		"comma separated should be valid sampling int value")
	assert.IsEqual(t, 456, getFirstIntValue("abc 456"),
		"space separated should be valid sampling int value")
	assert.IsEqual(t, 789, getFirstIntValue("abc 789 KB"),
		"extra chars should be valid sampling int value")
}

func TestGetFirstStringValue(t *testing.T) {
	assert.IsEqual(t, invalidSamplingStringValue, getFirstStringValue(""),
		"empty line should be invalid sampling string value")
	assert.IsEqual(t, invalidSamplingStringValue, getFirstStringValue(" "),
		"whitespace line should be invalid sampling string value")

	assert.IsEqual(t, "abc", getFirstStringValue("abc:123"),
		"comma separated should be valid sampling string value")
	assert.IsEqual(t, "xyz", getFirstStringValue("xyz 456"),
		"space separated should be valid sampling string value")
	assert.IsEqual(t, "uxa", getFirstStringValue("uxa 789 KB"),
		"extra chars should be valid sampling string value")
}

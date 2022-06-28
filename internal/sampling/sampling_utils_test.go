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

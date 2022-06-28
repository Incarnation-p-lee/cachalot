package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
)

func TestSampleThreadStat(t *testing.T) {
	testPID := 1
	threadStat := sampleThreadStat(testPID)

	assert.IsTrue(t, threadStat.ThreadsCount > 0, "thread count should be greater than 0")
}

func TestSampleThreadStatInvalidCount(t *testing.T) {
	testPID := 10000000
	threadStat := sampleThreadStat(testPID)

	assert.IsEqual(t, invalidThreadsCount, threadStat.ThreadsCount,
		"thread count should be invalid")
}

func TestGetThreadsCount(t *testing.T) {
	assert.IsEqual(t, 1, getThreadsCount("threads: 1"),
		"should have the same threads count")

	assert.IsEqual(t, invalidThreadsCount, getThreadsCount(""),
		"should be the invalid threads count")
}

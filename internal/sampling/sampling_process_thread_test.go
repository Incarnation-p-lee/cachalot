package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"testing"
)

func TestSampleThreadStat(t *testing.T) {
	threadsStatChan := make(chan snapshot.ThreadsStat)
	defer close(threadsStatChan)

	go sampleThreadsStat(1, threadsStatChan)
	threadsStat := <- threadsStatChan

	assert.IsTrue(t, threadsStat.ThreadsCount > 0, "thread count should be greater than 0")
}

func TestSampleThreadStatInvalidCount(t *testing.T) {
	threadsStatChan := make(chan snapshot.ThreadsStat)
	defer close(threadsStatChan)

	go sampleThreadsStat(100000000, threadsStatChan)
	threadsStat := <- threadsStatChan

	assert.IsEqual(t, invalidThreadsCount, threadsStat.ThreadsCount,
		"thread count should be invalid")
}

func TestGetThreadsCount(t *testing.T) {
	assert.IsEqual(t, 1, getThreadsCount("threads: 1"),
		"should have the same threads count")

	assert.IsEqual(t, invalidThreadsCount, getThreadsCount(""),
		"should be the invalid threads count")
}

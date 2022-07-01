package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"testing"
)

func TestInitTotalMemoryInKB(t *testing.T) {
	initTotalMemoryInKB()

	assert.IsTrue(t, totalMemoryInKB != invalidMemoryInKB,
		"should not be invalid total memory in KB")
	assert.IsTrue(t, totalMemoryInKB > 0,
		"total memory in KB should be greater than 0")
}

func TestGetMemoryInKB(t *testing.T) {
	assert.IsEqual(t, 123, getMemoryInKB("totalMemory: 123"),
		"should have the same total memory")

	assert.IsEqual(t, invalidMemoryInKB, getMemoryInKB(""),
		"should be the invalid memory in KB")
}

func TestGetMemoryName(t *testing.T) {
	assert.IsEqual(t, "VmSize", getMemoryName("VmSize: 123"),
		"should have the same memory name")

	assert.IsEqual(t, invalidMemoryName, getMemoryName(""),
		"should be the invalid memory name")
}

func TestSampleMemoryStat(t *testing.T) {
	memoryStatChan := make(chan snapshot.MemoryStat)
	defer close(memoryStatChan)

	go sampleMemoryStat(1, memoryStatChan)
	memoryStat := <-memoryStatChan

	assert.IsTrue(t, memoryStat.VMSizeInKB != invalidMemoryInKB,
		"vmSize in KB should not be invalid")
	assert.IsTrue(t, memoryStat.VMRSSInKB != invalidMemoryInKB,
		"vmRSS in KB should not be invalid")
	assert.IsTrue(t, memoryStat.VMStkInKB != invalidMemoryInKB,
		"vmStk in KB should not be invalid")
}

func TestSampleMemoryStatInvalid(t *testing.T) {
	memoryStatChan := make(chan snapshot.MemoryStat)
	defer close(memoryStatChan)

	go sampleMemoryStat(10000000, memoryStatChan)
	memoryStat := <-memoryStatChan

	assert.IsTrue(t, memoryStat.VMSizeInKB == invalidMemoryInKB,
		"vmSize in KB should be invalid")
	assert.IsTrue(t, memoryStat.VMRSSInKB == invalidMemoryInKB,
		"vmRSS in KB should be invalid")
	assert.IsTrue(t, memoryStat.VMStkInKB == invalidMemoryInKB,
		"vmStk in KB should be invalid")
}

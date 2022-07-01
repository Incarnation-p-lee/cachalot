package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"testing"
)

func TestSampleCPUStat(t *testing.T) {
	cpuStatChan := make(chan snapshot.CPUStat)
	defer close(cpuStatChan)

	go sampleCPUStat(1, cpuStatChan)
	cpuStat := <-cpuStatChan

	assert.IsTrue(t, cpuStat.JiffiesInTotal != invalidJiffies,
		"cpu total jiffies should not be invalid.")
	assert.IsTrue(t, cpuStat.JiffiesInTotal != defaultJiffies,
		"cpu total jiffies should not be default value.")

	assert.IsTrue(t, cpuStat.UsageInPercentage <= 100.0,
		"the cpu usage percentage should be less than or equal to 100.0")
}

func TestSampleCPUJiffies(t *testing.T) {
	testPID := 1
	totalJiffies, processJiffies := sampleCPUJiffies(testPID)

	assert.IsTrue(t, processJiffies <= totalJiffies,
		"the process jiffies should be less than or eual to total jiffies")

	assert.IsTrue(t, processJiffies != invalidJiffies,
		"the process used jiffies should not be invalid.")

	assert.IsTrue(t, processJiffies != defaultJiffies,
		"the process used jiffies should not be default value.")
}

func TestTestSampleCPUJiffiesInvalidPID(t *testing.T) {
	testPID := 10000000
	totalJiffies, processJiffies := sampleCPUJiffies(testPID)

	assert.IsTrue(t, processJiffies <= totalJiffies,
		"the process jiffies should be less than or eual to total jiffies")

	assert.IsTrue(t, processJiffies == invalidJiffies,
		"the invalid process id used jiffies should be invalid.")

	assert.IsTrue(t, processJiffies != defaultJiffies,
		"the invalid process id used jiffies should not be default value.")
}

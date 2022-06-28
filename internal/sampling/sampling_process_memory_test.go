package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
)

func TestInitTotalMemoryInKB(t *testing.T) {
	assert.IsEqual(t, invalidMemoryInKB, totalMemoryInKB,
		"total memory in KB should be invalid")

	initTotalMemoryInKB()

	assert.IsTrue(t, totalMemoryInKB != invalidMemoryInKB,
		"should not be invalid total memory in KB")
	assert.IsTrue(t, totalMemoryInKB > 0,
		"total memory in KB should be greater than 0")
}

func TestGetTotalMemoryInKB(t *testing.T) {
	assert.IsEqual(t, 123, getTotalMemoryInKB("totalMemory: 123"),
		"should have the same total memory")

	assert.IsEqual(t, invalidMemoryInKB, getTotalMemoryInKB(""),
		"should be the invalid memory in KB")
}

package sampling

import (
    "testing"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestInitTotalMemoryInKB(t *testing.T) {
	assert.IsEqual(t, invalidMemoryInKB, totalMemoryInKB,
		"total memory in KB should be invalid" )

	initTotalMemoryInKB()

	assert.IsTrue(t, totalMemoryInKB != invalidMemoryInKB,
		"should not be invalid total memory in KB" )
	assert.IsTrue(t, totalMemoryInKB > 0,
		"total memory in KB should be greater than 0" )
}


package print

import (
    "time"
    "testing"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func TestPrintSnapshot(t *testing.T) {
    ops := options.CreateOptions()
    testProcesses := []snapshot.Process {
        snapshot.Process {},
    }
    testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)

    assert.IsNil(t, PrintSnapshot(testSnapshot, "", ops), "print snapshot should have nil error")
    assert.IsNil(t, PrintSnapshot(testSnapshot, "abc", ops), "print snapshot should have nil error")
    assert.IsNotNil(t, PrintSnapshot(testSnapshot, "", nil), "print snapshot should have error")
}


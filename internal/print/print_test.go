package print

import (
    "time"
    "testing"
    "internal/options"
    "internal/cmdflags"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func TestPrintSnapshotDefault(t *testing.T) {
    ops := options.CreateOptions()
    testProcesses := []snapshot.Process {
        snapshot.Process {},
    }
    testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)

    assert.IsNil(t, Snapshot(testSnapshot, "", ops), "print empty title snapshot should have nil error")
    assert.IsNil(t, Snapshot(testSnapshot, "abc", ops), "print text title snapshot should have nil error")
    assert.IsNotNil(t, Snapshot(testSnapshot, "", nil), "print snapshot should have error")

    cmdflags.ParseOptions(ops)
    assert.IsNil(t, Snapshot(testSnapshot, "", ops), "print parsed snapshot should have nil error")
}

func TestPrintTextSnapshot(t *testing.T) {
    ops := options.CreateOptions()
    testProcesses := []snapshot.Process {
        snapshot.Process {},
    }

    testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)
    ops.AppendOption(options.Option {
        Key: options.OutputFormat,
        Val: options.TextOutput,
    })

    assert.IsNil(t, Snapshot(testSnapshot, "", ops), "print text snapshot should have nil error")
}

func TestPrintJsonSnapshot(t *testing.T) {
    ops := options.CreateOptions()
    testProcesses := []snapshot.Process {
        snapshot.Process {},
    }

    testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)
    ops.AppendOption(options.Option {
        Key: options.OutputFormat,
        Val: options.JSONOutput,
    })

    assert.IsNil(t, Snapshot(testSnapshot, "", ops), "print json snapshot should have nil error")
}


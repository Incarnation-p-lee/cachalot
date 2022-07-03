package print

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/cmdflags"
	"internal/options"
	"strconv"
	"testing"
	"time"
)

func TestPrintSnapshotDefault(t *testing.T) {
	ops := options.CreateOptions()
	testProcesses := []snapshot.Process{
		snapshot.Process{},
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
	testProcesses := []snapshot.Process{
		snapshot.Process{},
	}

	testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)
	ops.AppendOption(options.Option{
		Key: options.OutputFormat,
		Val: options.TextOutput,
	})

	assert.IsNil(t, Snapshot(testSnapshot, "", ops), "print text snapshot should have nil error")
}

func TestPrintJsonSnapshot(t *testing.T) {
	ops := options.CreateOptions()
	testProcesses := []snapshot.Process{
		snapshot.Process{},
	}

	testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)
	ops.AppendOption(options.Option{
		Key: options.OutputFormat,
		Val: options.JSONOutput,
	})

	assert.IsNil(t, Snapshot(testSnapshot, "", ops), "print json snapshot should have nil error")
}

func TestReconcileSnapshot(t *testing.T) {
	ops, topCount := options.CreateOptions(), 1
	ops.AppendOption(options.Option{
		Key: options.TopCount,
		Val: strconv.Itoa(topCount),
	})

	testProcesses := []snapshot.Process{
		snapshot.Process{},
		snapshot.Process{},
		snapshot.Process{},
	}

	testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)
	reconcileSnapshot(&testSnapshot, ops)

	assert.IsEqual(t, topCount, len(testSnapshot.Processes),
		"process count after reconcile should be the same as top count")
}

func TestReconcileSnapshotSortedByCPU(t *testing.T) {
	testProcesses := []snapshot.Process{
		snapshot.Process{
			CPUStat: snapshot.CPUStat{UsageInPercentage: 12.0},
		},
		snapshot.Process{
			CPUStat: snapshot.CPUStat{UsageInPercentage: 21.0},
		},
		snapshot.Process{
			CPUStat: snapshot.CPUStat{UsageInPercentage: 32.0},
		},
		snapshot.Process{
			CPUStat: snapshot.CPUStat{UsageInPercentage: 21.0},
		},
	}

	testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)
	reconcileSnapshotSortedBy(&testSnapshot, "cpu")

	for i := 1; i < len(testSnapshot.Processes); i++ {
		first, second := testSnapshot.Processes[i-1], testSnapshot.Processes[i]
		assert.IsTrue(t, first.CPUStat.UsageInPercentage >= second.CPUStat.UsageInPercentage,
			"the processes should be sorted by cpu in desc order")
	}
}

func TestReconcileSnapshotSortedByMemory(t *testing.T) {
	testProcesses := []snapshot.Process{
		snapshot.Process{
			MemoryStat: snapshot.MemoryStat{UsageInPercentage: 2.0},
		},
		snapshot.Process{
			MemoryStat: snapshot.MemoryStat{UsageInPercentage: 11.0},
		},
		snapshot.Process{
			MemoryStat: snapshot.MemoryStat{UsageInPercentage: 22.0},
		},
		snapshot.Process{
			MemoryStat: snapshot.MemoryStat{UsageInPercentage: 11.0},
		},
	}

	testSnapshot := snapshot.CreateSnapshot(time.Now(), testProcesses)
	reconcileSnapshotSortedBy(&testSnapshot, "memory")

	for i := 1; i < len(testSnapshot.Processes); i++ {
		first, second := testSnapshot.Processes[i-1], testSnapshot.Processes[i]
		assert.IsTrue(t, first.MemoryStat.UsageInPercentage >= second.MemoryStat.UsageInPercentage,
			"the processes should be sorted by memory in desc order")
	}
}

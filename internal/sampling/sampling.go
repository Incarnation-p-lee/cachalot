package sampling

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
	"internal/options"
	"time"
)

func init() {
	initTotalMemoryInKB()
}

func Sample(ops *options.Options) snapshot.Snapshot {
	if ops == nil {
		return snapshot.Snapshot{}
	}

	snapshot := snapshot.Snapshot{
		Timestamp: time.Now(),
		Network:   sampleNetwork(),
		Processes: []snapshot.Process{},
	}

	processes := sampleProcesses(ops, snapshot)

	snapshot.AppendProcesses(processes)

	return snapshot
}

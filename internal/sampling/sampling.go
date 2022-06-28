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

	timestamp, processes := time.Now(), sampleAllProcesses(ops)

	return snapshot.CreateSnapshot(timestamp, processes)
}

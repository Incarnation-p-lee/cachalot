package sampling

import (
    "time"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func Sample(ops *options.Options) snapshot.Snapshot {
    if ops == nil {
        return snapshot.Snapshot {}
    }

    snapshot := snapshot.CreateSnapshot(time.Now())
    snapshot.Processes = sampleAllProcess(ops)

    return snapshot
}


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

    timestamp, processes := time.Now(), sampleAllProcess(ops)

    return snapshot.CreateSnapshot(timestamp, processes)
}


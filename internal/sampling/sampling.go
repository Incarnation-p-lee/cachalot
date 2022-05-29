package sampling

import (
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/snapshot"
)

func Sample(ops *options.Options) snapshot.Snapshot {
    if ops == nil {
        return snapshot.Snapshot {}
    }

    return sampleSnapshot()
}

func sampleSnapshot() snapshot.Snapshot {
    return snapshot.Snapshot {}
}


package sampling

import (
    "testing"
    "internal/options"
    "github.com/Incarnation-p-lee/cachalot/pkg/assert"
)

func TestSample(t *testing.T) {
    ops := options.CreateOptions()

    assert.IsNotNil(t, Sample(nil), "Nil options will have non nil snapshot.")
    assert.IsNotNil(t, Sample(ops), "Non nil options will have non nil snapshot.")
}


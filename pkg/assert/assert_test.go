package assert

import (
    "testing"
)

func TestIsEqual(t *testing.T) {
    IsEqual(t, "abc", "abc", "test assert same string")
    IsEqual(t, 1, 1, "test assert same int")
    IsEqual(t, true, true, "test assert same bool")
    IsEqual(t, 4.0, 4.0, "test assert same float")
}

func TestIsTrue(t *testing.T) {
    IsTrue(t, true, "test assert true")
}

func TestIsFalse(t *testing.T) {
    IsFalse(t, false, "test assert false")
}


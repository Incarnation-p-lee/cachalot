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

func TestIsNil(t *testing.T) {
    IsNil(t, nil, "test assert is nil")
}

func TestIsNotNil(t *testing.T) {
    IsNotNil(t, map[string]int {}, "test assert is not nil")
    IsNotNil(t, []string {}, "test assert is not nil")
}


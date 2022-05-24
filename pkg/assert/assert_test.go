package assert

import (
    "testing"
)

func TestAssertEqual(t *testing.T) {
    AssertEqual(t, "abc", "abc", "test assert same string")
    AssertEqual(t, 1, 1, "test assert same int")
    AssertEqual(t, true, true, "test assert same bool")
    AssertEqual(t, 4.0, 4.0, "test assert same float")
}

func TestAssertTrue(t *testing.T) {
    AssertTrue(t, true, "test assert true")
}

func TestAssertFalse(t *testing.T) {
    AssertFalse(t, false, "test assert false")
}


package assert

import (
    "fmt"
    "testing"
)

func logAndFail(t *testing.T, message string) {
    t.Logf(message)
    t.Fail()
}

// IsEqual will check actual is the same as expect, or it will fail the test.
func IsEqual(t *testing.T, expect, actual interface{}, message string) {
    if expect == actual {
        return
    }

    msg := fmt.Sprintf("%s, assert equal failed, expect %+v != actual %+v", message, expect, actual)

    logAndFail(t, msg)
}

// IsTrue will check actual is true, or it will fail the test.
func IsTrue(t *testing.T, actual bool, message string) {
    if actual {
        return
    }

    msg := fmt.Sprintf("%s, assert true failed, actual %+v", message, actual)

    logAndFail(t, msg)
}

// IsNil will check actual is nil, or it will fail the test.
func IsNil(t *testing.T, actual interface{}, message string) {
    if actual == nil {
        return
    }

    msg := fmt.Sprintf("%s, assert is nil failed, actual %+v", message, actual)

    logAndFail(t, msg)
}

// IsNotNil will check actual is not nil, or it will fail the test.
func IsNotNil(t *testing.T, actual interface{}, message string) {
    if actual != nil {
        return
    }

    msg := fmt.Sprintf("%s, assert is not nil failed, actual %+v", message, actual)

    logAndFail(t, msg)
}

// IsFalse will check actual is false, or it will fail the test.
func IsFalse(t *testing.T, actual bool, message string) {
    if !actual {
        return
    }

    msg := fmt.Sprintf("%s, assert false failed, actual %+v", message, actual)

    logAndFail(t, msg)
}


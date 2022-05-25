package assert

import (
    "fmt"
    "testing"
)

func logAndFail(t *testing.T, message string) {
    t.Logf(message)
    t.Fail()
}

// AssertFalse will check actual is the same as expect, or it will fail the test.
func AssertEqual(t *testing.T, expect, actual interface{}, message string) {
    if expect == actual {
        return
    }

    msg := message

    if len(msg) == 0 {
        msg = fmt.Sprintf("assert equal failed, expect %+v != actual %+v",
            expect, actual)
    }

    logAndFail(t, msg)
}

// AssertFalse will check actual is true, or it will fail the test.
func AssertTrue(t *testing.T, actual bool, message string) {
    if actual {
        return
    }

    msg := message

    if len(msg) == 0 {
        msg = fmt.Sprintf("assert true failed, actual %+v", actual)
    }

    logAndFail(t, msg)
}

// AssertFalse will check actual is false, or it will fail the test.
func AssertFalse(t *testing.T, actual bool, message string) {
    if !actual {
        return
    }

    msg := message

    if len(msg) == 0 {
        msg = fmt.Sprintf("assert false failed, actual %+v", actual)
    }

    logAndFail(t, msg)
}


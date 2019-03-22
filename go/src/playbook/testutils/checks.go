package testutils

import "testing"

func Equal(t *testing.T, expected, actual interface{}, msg string, args ...interface{}) {
	if expected != actual {
		t.Errorf(msg, args...)
	}
}

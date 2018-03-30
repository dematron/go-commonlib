package commonlib

import (
	"testing"
	"time"
)

type testpair struct {
	value   int64
	correct bool
}

var tests = []testpair{
	{60, false},
	{time.Now().Unix(), true},
}

func TestMTimeCheck(t *testing.T) {
	testfile := "LICENSE"
	for _, pair := range tests {
		v := MTimeCheck(testfile, pair.value)
		if v != pair.correct {
			t.Errorf("For %d, expected %t, got %t.", pair.value, pair.correct, v)
		}
	}
}

package humanize

import (
	"testing"
)

type testpair struct {
	value uint64
	correct string
}

var tests = []testpair{
	{ 120, "120.00 " },
	{ 12300, "12.30 K" },
	{ 12340000, "12.34 M" },
	{ 123450000000, "123.45 G" },
	{ 12345600000000, "12.35 T" },
	{ 123456700000000000, "123.46 P" },
	{ 10000000000000000000, "10.00 E" },
}

func TestBytes2Human(t *testing.T) {
	for _, pair := range tests {
		v := Bytes2Human(pair.value, "")
		if v != pair.correct {
			t.Errorf("For %d, expected %q, got %q.", pair.value, pair.correct, v)
		}
	}
}


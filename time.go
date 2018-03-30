package commonlib

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Time
type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := t.Time.Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	t.Time = time.Unix(int64(ts), 0)

	return nil
}

// MTimeCheck get modify time of file "name" and check if it more or less then "seconds"
// return true if modify time is less or equal then "seconds", else false
func MTimeCheck(name string, seconds int64) bool {
	sec := time.Duration(seconds)
	needtime := sec * time.Second
	now := time.Now()

	fi, err := os.Stat(name)
	if e := CheckError(err, ErrorTrace, ErrorTracepath); e != nil {
		return false
	}

	mtime := fi.ModTime()
	diff := now.Sub(mtime)

	// Modify time is less or equal then "seconds"
	if diff <= (needtime) {
		return true
	}

	return false
}

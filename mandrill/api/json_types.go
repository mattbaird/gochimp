package api

import (
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	l := len(s)
	switch {
	case l == 12:
		t.Time, err = time.Parse(`"2006-01-02"`, s)
	case l == 21:
		t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, s)
	case l == 27:
		t.Time, err = time.Parse(`"2006-01-02 15:04:05.00000"`, s)
	case l == 9:
		t.Time, err = time.Parse(`"2006-01"`, s)
	}
	return
}

//APITimeFormat is the format string for time.Format
const APITimeFormat = "2006-01-02 15:04:05"

// nolint: deadcode, varcheck, unused
func apiTime(t interface{}) interface{} {
	switch ti := t.(type) {
	case time.Time:
		return ti.Format(APITimeFormat)
	case string:
		return ti
	}
	return t
}

// TS is a timestamp format for mandrill json
type TS struct {
	time.Time
}

func (t *TS) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	*t = TS{time.Unix(i, 0)}

	return nil
}

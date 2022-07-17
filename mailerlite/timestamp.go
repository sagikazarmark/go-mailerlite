package mailerlite

import (
	"strconv"
	"time"
)

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an RFC3339 or Unix timestamp.
// All exported methods of time.Time can be called on Timestamp.
type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in RFC3339 or Unix format.
func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		t.Time = time.Unix(i, 0)
		if t.Time.Year() > 3000 {
			t.Time = time.Unix(0, i*1e6)
		}
	} else {
		t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, str)
	}
	return
}

// Equal reports whether t and u are equal based on time.Equal.
func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time.Equal(u.Time)
}

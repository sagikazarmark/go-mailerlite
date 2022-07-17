package mailerlite

import (
	"encoding/json"
	"strconv"
)

// WeakInt can be used in places where the return type may be both int and string.
type WeakInt int

// UnmarshalJSON implements the json.Unmarshaler interface.
func (w *WeakInt) UnmarshalJSON(data []byte) (err error) {
	var i int
	err = json.Unmarshal(data, &i)
	if err == nil {
		*w = WeakInt(i)

		return
	}

	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return
	}

	i64, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		*w = WeakInt(i64)
	}

	return
}

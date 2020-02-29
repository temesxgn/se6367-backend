package scalar

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"time"

	gqlgen "github.com/99designs/gqlgen/graphql"
)

// MarshalTimestamp -
func MarshalTimestampScalar(t time.Time) gqlgen.Marshaler {
	return gqlgen.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

// UnmarshalTimestamp -
func UnmarshalTimestampScalar(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(json.Number); ok {
		if t, err := tmpStr.Int64(); err == nil {
			return time.Unix(t, 0), nil
		}
	}

	return time.Time{}, errors.New("time should be a unix timestamp")
}

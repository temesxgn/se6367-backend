package scalar

import (
	"encoding/json"
	"fmt"
	gqlgen "github.com/99designs/gqlgen/graphql"
	"io"
)

// MarshalMapScalar - marshal map
func MarshalMapScalar(val map[string]interface{}) gqlgen.Marshaler {
	return gqlgen.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

// MarshalMapScalar - unmarshal map
func UnmarshalMapScalar(v interface{}) (map[string]interface{}, error) {
	if m, ok := v.(map[string]interface{}); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}

package jsonutils

// TODO move out to seperate repo

import (
	"bytes"
	"encoding/json"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/tdewolff/minify"
	minjson "github.com/tdewolff/minify/json"
)

// Unmarshal - converts the json string to the data struct using the standard json encoding one
func Unmarshal(jsonString string, data interface{}) error {
	return json.Unmarshal([]byte(jsonString), &data)
}

// UnmarshalPB - converts the json string to the data struct using the proto unmarshal
func UnmarshalPB(jsonString string, pb proto.Message) error {
	return jsonpb.UnmarshalString(jsonString, pb)
}

// Marshal - converts the struct to a json string using standard json marshall
func Marshal(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

// MarshalPB - converts the struct to a json string using proto marshaller
func MarshalPB(pb proto.Message) (string, error) {
	marshaller := &jsonpb.Marshaler{OrigName: true, Indent: "  "}
	return marshaller.MarshalToString(pb)
}

// PrettyPrint - pretty prints a json string
func PrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

// Minify - minifies the json string
func Minify(jsonStr string) string {
	m := minify.New()
	r := bytes.NewBufferString(jsonStr)
	w := &bytes.Buffer{}
	if err := minjson.Minify(m, w, r, nil); err != nil {
		return jsonStr
	}
	return w.String()
}

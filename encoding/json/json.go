package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func SafeMarshal(v interface{}) []byte {
	bs, _ := Marshal(v)
	return bs
}

func SafeMarshalString(v interface{}) string {
	bs := SafeMarshal(v)
	return string(bs)
}

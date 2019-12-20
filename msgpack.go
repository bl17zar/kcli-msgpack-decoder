package main

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack"
)

// Msgpack is an example of how to create a plugin
// to decode msgpack encoded kafka messages.
// To compile:
//     go build -buildmode=plugin -o msgpack.so msgpack.go
// Then start kcli like:
//     kcli -d ./msgpack.so
type Msgpack struct{}

// Decode is required in order to be a plugin
func (m Msgpack) Decode(topic string, b []byte) ([]byte, error) {
	if topic == "" {
		return b, nil
	}

	var out map[string]interface{}

	err := msgpack.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}

	repr, err := json.Marshal(out)

	if err != nil {
		return b, nil
	}

	return repr, nil
}

// Decoder is the symbol that kcli will search for when
// using this as a plugin.
var Decoder Msgpack

func main() {}

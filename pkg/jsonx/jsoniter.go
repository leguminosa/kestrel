package jsonx

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

type (
	// jsonIterator implements Client interface
	jsonIterator struct {
		client jsoniter.API
	}
)

func newJSONIterator() *jsonIterator {
	return &jsonIterator{
		client: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

func (c *jsonIterator) Marshal(v interface{}) ([]byte, error) {
	return c.client.Marshal(v)
}

func (c *jsonIterator) Unmarshal(data []byte, v interface{}) error {
	return c.client.Unmarshal(data, v)
}

func (c *jsonIterator) Encode(writer io.Writer, v interface{}) error {
	return c.client.NewEncoder(writer).Encode(v)
}

func (c *jsonIterator) Decode(reader io.Reader, v interface{}) error {
	return c.client.NewDecoder(reader).Decode(v)
}

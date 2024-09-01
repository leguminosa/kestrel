package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want int
	}{
		{
			name: "int",
			v:    1,
			want: 1,
		},
		{
			name: "int32",
			v:    int32(32),
			want: 32,
		},
		{
			name: "int64",
			v:    int64(64),
			want: 64,
		},
		{
			name: "float32",
			v:    float32(32.8),
			want: 32,
		},
		{
			name: "float64",
			v:    float64(64.1),
			want: 64,
		},
		{
			name: "string",
			v:    " 911  ",
			want: 911,
		},
		{
			name: "unparseable string",
			v:    " 91 2  ",
			want: 0,
		},
		{
			name: "slice of byte",
			v:    []byte("52 "),
			want: 52,
		},
		{
			name: "invalid slice of byte",
			v:    []byte("invalid string"),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToInt(tt.v)
			assert.Equal(t, tt.want, got)
		})
	}
}

package response

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		want    string
		wantErr bool
	}{
		{
			name: "success response",
			i: map[string]interface{}{
				"key": "value",
			},
			want:    "{\"data\":{\"key\":\"value\"},\"message\":\"OK\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := OK(c, tt.i)
			if !assert.Equal(t, tt.wantErr, err != nil) {
				return
			}

			got := c.(*mockEchoContext).getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestBadRequest(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		message string
		want    string
		wantErr bool
	}{
		{
			name:    "bad request response",
			err:     errors.New("mocked error"),
			message: "mocked message",
			want:    "{\"error\":\"mocked error\",\"message\":\"mocked message\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := BadRequest(c, tt.err, tt.message)
			if !assert.Equal(t, tt.wantErr, err != nil) {
				return
			}

			got := c.(*mockEchoContext).getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestInternalServerError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		message string
		want    string
		wantErr bool
	}{
		{
			name:    "server error response",
			err:     errors.New("mocked error"),
			message: "mocked message",
			want:    "{\"error\":\"mocked error\",\"message\":\"mocked message\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InternalServerError(c, tt.err, tt.message)
			if !assert.Equal(t, tt.wantErr, err != nil) {
				return
			}

			got := c.(*mockEchoContext).getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func Test_buildErrorResponse(t *testing.T) {
	tests := []struct {
		name    string
		code    int
		err     error
		message string
		want    string
		wantErr bool
	}{
		{
			name:    "server error response",
			code:    500,
			err:     errors.New("mocked error"),
			message: "mocked message",
			want:    "{\"error\":\"mocked error\",\"message\":\"mocked message\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := buildErrorResponse(c, tt.code, tt.err, tt.message)
			if !assert.Equal(t, tt.wantErr, err != nil) {
				return
			}

			got := c.(*mockEchoContext).getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func Test_buildJSONResponse(t *testing.T) {
	tests := []struct {
		name    string
		code    int
		i       interface{}
		want    string
		wantErr bool
	}{
		{
			name: "success response",
			code: 200,
			i: map[string]interface{}{
				"message": "ok",
			},
			want:    "{\"message\":\"ok\"}\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := buildJSONResponse(c, tt.code, tt.i)
			if !assert.Equal(t, tt.wantErr, err != nil) {
				return
			}

			got := c.(*mockEchoContext).getResponseBody()
			assert.Equal(t, tt.want, string(got))
		})
	}
}

package xcontext

import (
	"context"
	"net/http"
)

type contextKey int

const (
	ctxKeyUserID contextKey = iota
)

var (
	ctxMap = map[contextKey]func(ctx context.Context, r *http.Request) context.Context{
		ctxKeyUserID: userIDToContext,
	}
)

func userIDToContext(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, ctxKeyUserID, r.Header.Get("UserId"))
}

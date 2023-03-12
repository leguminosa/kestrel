package xcontext

import (
	"context"
	"net/http"
)

func GetAllContextFromIncomingRequest(r *http.Request) context.Context {
	return populate(r, getContextList(r)...)
}

func getContextList(r *http.Request) []contextKey {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" && contentType != "application/vnd.api+json" {
		return []contextKey{}
	}

	return []contextKey{
		ctxKeyUserID,
	}
}

func populate(r *http.Request, opts ...contextKey) context.Context {
	ctx := r.Context()

	for _, opt := range opts {
		opp, found := ctxMap[opt]
		if found {
			ctx = opp(ctx, r)
		}
	}

	return ctx
}

package middleware

import (
	"net/http"

	"github.com/leguminosa/kestrel/pkg/util/xcontext"
)

func PopulateContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(xcontext.GetAllContextFromIncomingRequest(r))

		next.ServeHTTP(w, r)
	})
}

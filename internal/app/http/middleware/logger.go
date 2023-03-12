package middleware

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/leguminosa/kestrel/pkg/envx"
)

func DumpIncomingRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if envx.IsProduction() {
			next.ServeHTTP(w, r)
		}

		dumpReq, _ := httputil.DumpRequest(r, true)
		log.Printf("Incoming request:\n%s\n", dumpReq)

		next.ServeHTTP(w, r)
	})
}

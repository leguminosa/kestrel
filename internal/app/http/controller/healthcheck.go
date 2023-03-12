package controller

import (
	"net/http"

	"github.com/leguminosa/kestrel/internal/app/http/wrapper"
	"github.com/leguminosa/kestrel/pkg/envx"
)

func (ctrl *Controller) HealthCheck(w http.ResponseWriter, r *http.Request) {
	wrapper.OK(w, map[string]interface{}{
		"env":     envx.ServiceEnv(),
		"service": "kestrel",
		"status":  "ok",
	})
}

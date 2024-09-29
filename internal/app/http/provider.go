package http

import (
	"github.com/gorilla/mux"
	"github.com/leguminosa/kestrel/internal/app/http/controller"
	"github.com/leguminosa/kestrel/pkg/httpx"
)

type controllerGroup struct {
	character controller.Character
}

func provideHTTPServer(ctrl *controllerGroup) *mux.Router {
	router := httpx.NewRouter()
	v1 := router.PathPrefix("/v1").Subrouter()

	ctrl.character.RegisterRoutes(v1)

	return router.Mux()
}

package http

import (
	"github.com/gorilla/mux"
	"github.com/leguminosa/kestrel/internal/app/http/controller"
)

func register(handler *mux.Router, ctrl *controller.Controller) {
	handler.HandleFunc("/", ctrl.HealthCheck)
	handler.HandleFunc("/ping", ctrl.HealthCheck)

	handler.HandleFunc("/gear/setoption", ctrl.GearOptionRouter)
	handler.HandleFunc("/gear/setoption/{id}", ctrl.GearOptionRouterWithID)
}

package http

import (
	"github.com/leguminosa/kestrel/internal/app/http/controller"
	"github.com/leguminosa/kestrel/pkg/util/httpx"
)

func register(handler *httpx.Router, ctrl *controller.Controller) {
	handler.GET("/", ctrl.HealthCheck)
	handler.GET("/ping", ctrl.HealthCheck)

	// gear set option
	handler.GET("/setoption", ctrl.GetGearSetOptions)
	handler.GET("/setoption/{id}", ctrl.GetGearSetOption)
	handler.POST("/setoption", ctrl.CreateGearSetOption)
	handler.PUT("/setoption/{id}", ctrl.UpdateGearSetOption)
	handler.DELETE("/setoption/{id}", ctrl.DeleteGearSetOption)
}

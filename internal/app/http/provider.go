package http

import (
	"github.com/labstack/echo/v4"
	"github.com/leguminosa/kestrel/internal/app/http/instructions"
)

type controller struct {
	characterHandler instructions.CharacterServer
}

func provideHTTPServer(ctrl *controller) *echo.Echo {
	e := echo.New()

	e.GET("/characters/:id", ctrl.characterHandler.GetByID)

	return e
}

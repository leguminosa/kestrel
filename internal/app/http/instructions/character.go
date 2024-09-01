package instructions

import "github.com/labstack/echo/v4"

type CharacterServer interface {
	GetByID(c echo.Context) error
}

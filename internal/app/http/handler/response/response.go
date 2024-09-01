package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func OK(c echo.Context, i interface{}) error {
	return buildJSONResponse(c, http.StatusOK, map[string]interface{}{
		"message": "OK",
		"data":    i,
	})
}

func BadRequest(c echo.Context, err error, message string) error {
	return buildErrorResponse(c, http.StatusBadRequest, err, message)
}

func InternalServerError(c echo.Context, err error, message string) error {
	return buildErrorResponse(c, http.StatusInternalServerError, err, message)
}

func buildErrorResponse(c echo.Context, code int, err error, message string) error {
	content := map[string]interface{}{
		"message": message,
	}
	if err != nil {
		content["error"] = err.Error()
	}
	return buildJSONResponse(c, code, content)
}

func buildJSONResponse(c echo.Context, code int, i interface{}) error {
	return c.JSON(code, i)
}

package http

import (
	"github.com/labstack/echo/v4"
	"github.com/leguminosa/kestrel/internal/config"
)

type (
	serverImpl struct {
		port string
		echo *echo.Echo

		err error
	}
)

func newServer(
	cfg config.Config,
	e *echo.Echo,
) *serverImpl {
	return &serverImpl{
		port: cfg.Server.Port,
		echo: e,
	}
}

func newServerWithError(err error) *serverImpl {
	return &serverImpl{err: err}
}

func (s *serverImpl) Start() error {
	if s.err != nil {
		return s.err
	}
	return s.echo.Start(s.port)
}

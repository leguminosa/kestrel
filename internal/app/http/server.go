package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leguminosa/kestrel/internal/config"
)

type (
	serverImpl struct {
		server *http.Server

		err error
	}
)

func newServer(
	cfg config.ServerConfig,
	router *mux.Router,
) *serverImpl {
	return &serverImpl{
		server: &http.Server{
			Addr:    cfg.Port,
			Handler: router,
		},
	}
}

func newServerWithError(err error) *serverImpl {
	return &serverImpl{err: err}
}

func (s *serverImpl) Start() error {
	if s.err != nil {
		return s.err
	}
	return s.server.ListenAndServe()
}

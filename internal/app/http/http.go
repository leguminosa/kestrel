package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leguminosa/kestrel/internal/app/http/controller"
	"github.com/leguminosa/kestrel/internal/app/http/middleware"
	"github.com/leguminosa/kestrel/pkg/config"
	"github.com/leguminosa/kestrel/pkg/database"
)

func InitApp(cfg *config.Config, db *database.Database) Server {
	handler := mux.NewRouter()

	register(handler, controller.New(db))

	handler.Use(
		middleware.DumpIncomingRequest,
	)

	return &serverImpl{
		port:    cfg.Port,
		handler: handler,
	}
}

type (
	Server interface {
		Run() error
	}
	serverImpl struct {
		port    string
		handler http.Handler
	}
)

func (s *serverImpl) Run() error {
	serv := &http.Server{
		Addr:    s.port,
		Handler: s.handler,
	}
	return serv.ListenAndServe()
}

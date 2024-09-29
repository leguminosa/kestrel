package httpx

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	router := mux.NewRouter()
	return &Router{Router: router}
}

func wrapRouter(router *mux.Router) *Router {
	return &Router{Router: router}
}

func (r *Router) PathPrefix(tpl string) *Route {
	return wrapRoute(r.Router.PathPrefix(tpl))
}

func (r *Router) Mux() *mux.Router {
	return r.Router
}

func (r *Router) Get(path string, handler http.Handler) {
	r.Router.Handle(path, handler).Methods(http.MethodGet)
}

func (r *Router) Post(path string, handler http.Handler) {
	r.Router.Handle(path, handler).Methods(http.MethodPost)
}

func (r *Router) Put(path string, handler http.Handler) {
	r.Router.Handle(path, handler).Methods(http.MethodPut)
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.Router.Handle(path, handler).Methods(http.MethodDelete)
}

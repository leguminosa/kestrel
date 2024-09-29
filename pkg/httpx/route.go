package httpx

import "github.com/gorilla/mux"

type Route struct {
	*mux.Route
}

func wrapRoute(route *mux.Route) *Route {
	return &Route{Route: route}
}

func (r *Route) Subrouter() *Router {
	return wrapRouter(r.Route.Subrouter())
}

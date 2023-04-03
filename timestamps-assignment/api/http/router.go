package http

import (
	"github.com/gorilla/mux"
)

type Router struct {
	timestampsRoutes TimestampsRoutesInterface
}

func InitMainRouter(t TimestampsRoutesInterface) *Router {
	return &Router{
		timestampsRoutes: t,
	}
}

func (r *Router) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = r.timestampsRoutes.SetTimestampsRoutes(router)
	return router
}

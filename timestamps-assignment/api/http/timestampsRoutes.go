package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

type TimestampsRoutes struct {
	timestampsHandler TimestampDataHandler
}

func NewTimestampsRoutes(t TimestampDataHandler) TimestampsRoutes {
	return TimestampsRoutes{
		timestampsHandler: t,
	}
}

type TimestampsRoutesInterface interface {
	SetTimestampsRoutes(router *mux.Router) *mux.Router
}

func (r TimestampsRoutes) SetTimestampsRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/ptlist", r.timestampsHandler.TimestampsMatching).Methods(http.MethodGet)
	return router
}

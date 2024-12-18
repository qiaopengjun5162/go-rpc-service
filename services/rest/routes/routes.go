package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/qiaopengjun5162/go-rpc-service/services/rest/service"
)

type Routes struct {
	router *chi.Mux
	svc    service.Service
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(r *chi.Mux, svc service.Service) Routes {
	return Routes{
		router: r,
		svc:    svc,
	}
}

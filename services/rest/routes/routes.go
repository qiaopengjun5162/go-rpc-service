package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/qiaopengjun5162/go-rpc-service/services/rest/service"
)

type Routes struct {
	router *chi.Mux
	svc    service.Service
}

// NewRoutes creates a new Routes object with the given chi router and service.
//
// Parameters:
//   - r: The chi router to use for this Routes object.
//   - svc: The service to use for this Routes object.
//
// Returns:
//   - A new Routes object with the given router and service.
func NewRoutes(r *chi.Mux, svc service.Service) Routes {
	return Routes{
		router: r,
		svc:    svc,
	}
}

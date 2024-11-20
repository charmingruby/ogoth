package health

import (
	"github.com/charmingruby/ogoth/internal/health/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewRestHandler(r *chi.Mux) *endpoint.Endpont {
	return endpoint.NewEndpoint(r)
}

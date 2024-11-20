package endpoint

import "github.com/go-chi/chi/v5"

func NewEndpoint(router *chi.Mux) *Endpont {
	return &Endpont{
		router: router,
	}
}

type Endpont struct {
	router *chi.Mux
}

func (e *Endpont) Register() {
	e.router.Get("/health", e.healthHandler())
}

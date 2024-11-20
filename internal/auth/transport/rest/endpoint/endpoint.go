package endpoint

import (
	"github.com/charmingruby/ogoth/internal/auth/transport/rest/client"
	"github.com/go-chi/chi/v5"
)

func NewEndpoint(r *chi.Mux, client *client.GoogleOAuth2) *Endpont {
	return &Endpont{
		router:             r,
		googleOAuth2Client: client,
	}
}

type Endpont struct {
	router             *chi.Mux
	googleOAuth2Client *client.GoogleOAuth2
}

func (e *Endpont) Register() {
	e.router.Get("/auth/{provider}", e.signInHandler())
	e.router.Get("/auth/{provider}/callback", e.callbackHandler())
	e.router.Get("/auth/tokens", e.retrieveTokensHandler())
}

package endpoint

import (
	"github.com/charmingruby/ogoth/internal/auth/transport/rest/client"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func NewEndpoint(r *chi.Mux, client *client.GoogleOAuth2, store *sessions.CookieStore) *Endpont {
	return &Endpont{
		router:             r,
		googleOAuth2Client: client,
		store:              store,
	}
}

type Endpont struct {
	router             *chi.Mux
	googleOAuth2Client *client.GoogleOAuth2
	store              *sessions.CookieStore
}

func (e *Endpont) Register() {
	e.router.Get("/auth/{provider}", e.signInHandler())
	e.router.Get("/auth/{provider}/callback", e.callbackHandler())
	e.router.Get("/auth/tokens", e.retrieveTokensHandler())
}

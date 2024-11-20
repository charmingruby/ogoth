package auth

import (
	"github.com/charmingruby/ogoth/internal/auth/transport/rest/client"
	"github.com/charmingruby/ogoth/internal/auth/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func NewRestHandler(r *chi.Mux, googleOAuth2Client client.GoogleOAuth2, store *sessions.CookieStore) *endpoint.Endpont {
	return endpoint.NewEndpoint(r, &googleOAuth2Client, store)
}

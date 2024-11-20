package endpoint

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func (e *Endpont) signInHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		if provider != GOOGLE_PROVIDER {
			slog.Error("provider not supported")
			http.Error(w, "provider not supported", http.StatusBadRequest)
			return
		}

		URL, err := url.Parse(e.googleOAuth2Client.Config.AuthCodeURL(OAUTH_STATE))
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		authUrl := URL.String()

		http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
	}
}

package endpoint

import (
	"io"
	"log/slog"
	"net/http"
)

func (e *Endpont) callbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query()["state"][0]
		if state != OAUTH_STATE {
			slog.Error("state not match")
			http.Error(w, "state not match", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")

		token, err := e.googleOAuth2Client.Config.Exchange(r.Context(), code)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := http.Get(USERINFO_URL + token.AccessToken)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userData, err := io.ReadAll(res.Body)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(userData)
	}
}

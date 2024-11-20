package endpoint

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/charmingruby/ogoth/internal/auth/core/model"
	"github.com/charmingruby/ogoth/internal/auth/transport/rest/constant"
)

func (e *Endpont) callbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query()["state"][0]
		if state != constant.OAUTH_STATE {
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

		res, err := http.Get(constant.USERINFO_URL + token.AccessToken)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userJSON, err := parseGoogleUserData(res)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		session, err := e.store.Get(r, "user-session")
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
			return
		}

		session.Values["user"] = userJSON
		err = session.Save(r, w)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Unable to save session", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userJSON)
	}
}

func parseGoogleUserData(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	user, err := model.JSONToUserModel(body)
	if err != nil {
		return nil, err
	}

	slog.Info(fmt.Sprintf("user: %+v\n", user))

	json, err := model.UserModelToJSON(user)
	if err != nil {
		return nil, err
	}

	return json, nil
}

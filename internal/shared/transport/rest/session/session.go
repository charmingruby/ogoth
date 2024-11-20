package session

import "github.com/gorilla/sessions"

func New(key string, environment string) *sessions.CookieStore {
	isProd := environment == "prod"

	store := sessions.NewCookieStore([]byte(key))

	store.MaxAge(60 * 60 * 24 * 7 * 30) // 1 month

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	return store
}

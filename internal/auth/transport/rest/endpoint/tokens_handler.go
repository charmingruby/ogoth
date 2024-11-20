package endpoint

import "net/http"

func (e *Endpont) retrieveTokensHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ...
	}
}

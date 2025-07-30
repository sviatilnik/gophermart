package handlers

import (
	"net/http"
)

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("Hello World"))

		w.WriteHeader(http.StatusOK)
	}
}

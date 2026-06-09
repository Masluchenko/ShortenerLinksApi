package middleware

import (
	"fmt"
	"net/http"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		athedHeader := r.Header.Get("Authorization")
		next.ServeHTTP(w, r)
		fmt.Println()
	})
}

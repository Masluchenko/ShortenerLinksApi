package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		athedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(athedHeader, "Bearer ")
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}

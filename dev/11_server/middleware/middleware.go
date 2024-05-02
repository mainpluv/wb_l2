package middleware

import (
	"log"
	"net/http"
)

// для логгирования
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

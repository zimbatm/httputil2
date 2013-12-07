package httputil2

import (
	"net/http"
)

// Adds a unique ID in the X-Request-Id request header.
//
// To be used with the uuid lib for example
func IdHandler(h http.Handler, id <-chan string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-Request-Id", <-id)
		h.ServeHTTP(w, r)
	})
}

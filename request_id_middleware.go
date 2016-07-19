package httputil2

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

// Tags a request X-Request-ID header with a given ID from a RequestIDGenerator.
//
// To be used with the uuid lib for example
//
// If g is nil it will use a RandomIDGenerator(32)
func RequestIDMiddleware(g RequestIDGenerator) Middleware {
	if g == nil {
		g = RandomIDGenerator(32)
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set(HeaderXRequestID, g())
			h.ServeHTTP(w, r)
		})
	}
}

type RequestIDGenerator func() string

// To be used with the IdHandler
//
// size must be a power of 2
//
// may fail if the random pool is exhausted
func RandomIDGenerator(size int) RequestIDGenerator {
	b := make([]byte, size/2)
	return func() string {
		_, err := rand.Read(b)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%x", b)
	}
}

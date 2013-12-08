package httputil2

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

const (
	HeaderRequestId = "X-Request-Id"
)

// Tags a request X-Request-Id header with a given ID from the IdGenerator.
//
// To be used with the uuid lib for example
func IdHandler(h http.Handler, g IdGenerator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set(HeaderRequestId, g())
		h.ServeHTTP(w, r)
	})
}

type IdGenerator func() string

// To be used with the IdHandler
//
// size must be a power of 2
//
// may fail if the random pool is exhausted
func RandomGenerator(size int) IdGenerator {
	b := make([]byte, size/2)
	return func() string {
		_, err := rand.Read(b)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%x", b)
	}
}

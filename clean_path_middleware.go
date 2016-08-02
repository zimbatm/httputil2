package httputil2

import (
	"net/http"
	"path"
)

type cleanPathMiddleWare struct {
	h http.Handler
}

func CleanPathMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return &cleanPathMiddleWare{h}
	}
}

func (c *cleanPathMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	if cleanPath != r.URL.Path {
		w.Header().Set(HeaderLocation, cleanPath)
		w.WriteHeader(StatusTemporaryRedirect)
	} else {
		c.h.ServeHTTP(w, r)
	}
}

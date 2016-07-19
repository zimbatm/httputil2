package httputil2

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type basicAuthHandler struct {
	h     http.Handler
	realm string
	c     BasicAuthChecker
}

type BasicAuthChecker func(user string, pass string) bool

func BasicAuthHandler(h http.Handler, realm string, c BasicAuthChecker) Middleware {
	return func(h http.Handler) http.Handler {
		return &basicAuthHandler{h, realm, c}
	}
}

func (self *basicAuthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	header := req.Header.Get(HeaderAuthorization)

	f := strings.Fields(header)
	if len(f) == 2 && f[0] == "Basic" {
		if b, err := base64.StdEncoding.DecodeString(f[1]); err == nil {
			kv := strings.SplitN(string(b), ":", 2)
			user := kv[0]
			pass := kv[1]

			ok := self.c(user, pass)
			if ok {
				// Authorized
				self.h.ServeHTTP(w, req)
				return
			}
		}
	}

	// Not Authorized
	w.Header().Set(HeaderWWWAuthenticate, `Basic realm="`+self.realm+`"`)
	w.WriteHeader(http.StatusUnauthorized)
}

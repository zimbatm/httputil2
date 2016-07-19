package httputil2

import (
	"net/http"
	"strconv"
	"strings"
)

// http://www.w3.org/TR/cors/

// Inspired by rack-cors : https://github.com/cyu/rack-cors

// CORS-specific HTTP header extensions
const (
	HeaderOrigin                        = "Origin"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"
)

func init() {
	RequestHeaders = append(
		RequestHeaders,
		HeaderOrigin,
		HeaderAccessControlRequestMethod,
		HeaderAccessControlRequestHeaders,
	)

	ResponseHeaders = append(
		RequestHeaders,
		HeaderAccessControlAllowCredentials,
		HeaderAccessControlAllowMethods,
		HeaderAccessControlAllowOrigin,
		HeaderAccessControlAllowOrigin,
		HeaderAccessControlExposeHeaders,
		HeaderAccessControlMaxAge,
	)
}

func CORSHandler(origin string, maxAge int) Middleware {
	if len(origin) == 0 {
		origin = "*"
	}
	if maxAge <= 0 {
		maxAge = 1728000
	}
	return func(h http.Handler) http.Handler {
		return &corsHandler{h, origin, true, maxAge, nil, []string{"any"}, nil}
	}
}

type corsHandler struct {
	h           http.Handler
	origin      string
	credentials bool
	maxAge      int
	methods     []string
	headers     []string
	expose      []string
}

func (self *corsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get(HeaderOrigin); len(origin) > 0 && self.allowed(origin) {
		w.Header().Set(HeaderAccessControlAllowOrigin, req.Header.Get("Origin"))

		// TODO: Control the allowed methods
		if len(self.methods) == 0 {
			w.Header().Add(HeaderAccessControlAllowMethods, GET)
		} else {
			for _, method := range self.methods {
				w.Header().Add(HeaderAccessControlAllowMethods, method)
			}
		}

		for _, header := range self.expose {
			w.Header().Add(HeaderAccessControlExposeHeaders, strings.ToLower(header))
		}

		w.Header().Set(HeaderAccessControlMaxAge, strconv.Itoa(self.maxAge))

		if self.credentials {
			w.Header().Set(HeaderAccessControlAllowCredentials, "true")
		}

		// Short circuit if it's an OPTIONS method
		if req.Method == OPTIONS && len(req.Header.Get(HeaderAccessControlRequestMethod)) > 0 {
			// Preflight
			w.Header().Set(HeaderAccessControlAllowHeaders, req.Header.Get(HeaderAccessControlRequestHeaders))
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	self.h.ServeHTTP(w, req)
}

// TODO: Check the method, headers and stuff, make this method external to the
//       handler so arbitrary rules can be applied.
func (self *corsHandler) allowed(origin string) bool {
	return self.origin == origin || self.origin == "*"
}

package httputil2

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler
type MiddlewareList struct {
	l []Middleware
}

// Appends more middlewares to the stack
func (ml MiddlewareList) Use(ms ...Middleware) {
	ml.l = append(ml.l, ms...)
}

// Puts together all the middlewares together with the last one as the end-point
func (ml *MiddlewareList) Chain(h http.Handler) http.Handler {
	if ml.l == nil {
		return h
	}
	for i := len(ml.l) - 1; i >= 0; i-- {
		h = ml.l[i](h)
	}
	return h
}

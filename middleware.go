package httputil2

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// Deprecated, use httputil2.Chain instead
type MiddlewareList struct {
	Handler http.Handler
	l       []Middleware
}

// Appends more middlewares to the stack
func (ml *MiddlewareList) Use(ms ...Middleware) {
	ml.l = append(ml.l, ms...)
}

// Puts the list of middleware together, with the last one being the provided
// ml.Handler.
func (ml *MiddlewareList) Chain() http.Handler {
	h := ml.Handler
	if h == nil {
		return nil
	}
	if ml.l == nil {
		return h
	}
	for i := len(ml.l) - 1; i >= 0; i-- {
		h = ml.l[i](h)
	}
	return h
}

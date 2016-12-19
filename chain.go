package httputil2

import "net/http"

// A chain of middlewares
type Chain []Middleware

func NewChain(chain ...Middleware) Chain {
	return Chain(chain)
}

func (c Chain) Append(m ...Middleware) Chain {
	return append(c, m...)
}

func (c Chain) Handle(h http.Handler) http.Handler {
	for i := len(c) - 1; i >= 0; i-- {
		h = c[i](h)
	}
	return h
}

func (c Chain) HandleFunc(h http.HandlerFunc) http.Handler {
	return c.Handle(http.HandlerFunc(h))
}

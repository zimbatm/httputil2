package httputil2

// TODO: implement the http.Flusher, http.CloseNotifier and http.Hijacker interfaces

import (
	"log"
	"net/http"
)

type ErrorCallback func(error)

// Recovers when a panic happens on a request
func RecoverHandler(callback ErrorCallback) Middleware {
	if callback == nil {
		callback = DefaultCallback
	}

	return func(h http.Handler) http.Handler {
		return &recoverHandler{h, callback}
	}
}

func DefaultCallback(err error) {
	log.Printf("ERR:", err)
}

type recoverHandler struct {
	h        http.Handler
	callback ErrorCallback
}

func (self *recoverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w2 := &recoverResponseWriter{w, false}
	defer handleError(w2, self.callback)
	self.h.ServeHTTP(w2, r)
}

func handleError(w *recoverResponseWriter, callback ErrorCallback) {
	var err error
	if x := recover(); x != nil {
		err = x.(error)

		if !w.wroteHeader {
			w.WriteHeader(http.StatusInternalServerError)
		}

		callback(err)
	}
}

type recoverResponseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (self *recoverResponseWriter) WriteHeader(status int) {
	self.wroteHeader = true
	self.ResponseWriter.WriteHeader(status)
}

func (self *recoverResponseWriter) Write(data []byte) (int, error) {
	// We have to re-implement that logic that's in the http lib unfortunately
	if !self.wroteHeader {
		self.WriteHeader(http.StatusOK)
	}
	return self.ResponseWriter.Write(data)
}

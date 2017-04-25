package httputil2

// TODO: implement the http.Flusher, http.CloseNotifier and http.Hijacker interfaces

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

type ErrorCallback func(error)

// Recovers when a panic happens on a request
func RecoveryMiddleware(callback ErrorCallback) Middleware {
	if callback == nil {
		callback = DefaultCallback
	}

	return func(h http.Handler) http.Handler {
		return &recoveryHandler{h, callback}
	}
}

func DefaultCallback(err error) {
	log.Println("ERR:", err)
}

type recoveryHandler struct {
	h        http.Handler
	callback ErrorCallback
}

func (self *recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (self *recoverResponseWriter) Flush() {
	(self.ResponseWriter.(http.Flusher)).Flush()
}

func (self *recoverResponseWriter) CloseNotify() <-chan bool {
	return (self.ResponseWriter.(http.CloseNotifier)).CloseNotify()
}

func (self *recoverResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return (self.ResponseWriter.(http.Hijacker)).Hijack()
}

var _ http.Hijacker = new(recoverResponseWriter)
var _ http.ResponseWriter = new(recoverResponseWriter)
var _ http.Flusher = new(recoverResponseWriter)
var _ http.CloseNotifier = new(recoverResponseWriter)

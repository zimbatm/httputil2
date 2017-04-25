package httputil2

// TODO: implement the http.Hijacker interfaces

import (
	"bufio"
	"net"
	"net/http"
	"time"
)

type HTTPLogger interface {
	LogRequest(r *http.Request, start time.Time)
	LogResponse(r *http.Request, start time.Time, status int, bytes int)
}

func LogMiddleware(f HTTPLogger) Middleware {
	return func(h http.Handler) http.Handler {
		return &logHandler{h, f}
	}
}

type logHandler struct {
	h http.Handler
	l HTTPLogger
}

func (self *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	self.l.LogRequest(r, start)

	lw := &logResponseWriter{w, 0, 0}
	self.h.ServeHTTP(lw, r)

	self.l.LogResponse(r, start, lw.status, lw.bytes)
}

type logResponseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (self *logResponseWriter) WriteHeader(status int) {
	self.status = status
	self.ResponseWriter.WriteHeader(status)
}

func (self *logResponseWriter) Write(data []byte) (n int, err error) {
	// We have to re-implement that logic that's in the http lib unfortunately
	if self.status == 0 {
		self.WriteHeader(http.StatusOK)
	}
	n, err = self.ResponseWriter.Write(data)
	self.bytes += n
	return
}

func (self *logResponseWriter) Flush() {
	(self.ResponseWriter.(http.Flusher)).Flush()
}

func (self *logResponseWriter) CloseNotify() <-chan bool {
	return (self.ResponseWriter.(http.CloseNotifier)).CloseNotify()
}

func (self *logResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return (self.ResponseWriter.(http.Hijacker)).Hijack()
}

var _ http.ResponseWriter = new(logResponseWriter)
var _ http.Flusher = new(logResponseWriter)
var _ http.CloseNotifier = new(logResponseWriter)
var _ http.Hijacker = new(logResponseWriter)

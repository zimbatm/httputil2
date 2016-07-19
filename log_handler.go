package httputil2

// TODO: implement the http.Hijacker interfaces

import (
	"io"
	"net/http"
	"time"
)

type LogFormatter interface {
	RequestLog(r *http.Request, start time.Time) string
	ResponseLog(r *http.Request, start time.Time, status int, bytes int) string
}

func LogHandler(w io.Writer, f LogFormatter) Middleware {
	return func(h http.Handler) http.Handler {
		return &logHandler{h, w, f}
	}
}

type logHandler struct {
	h http.Handler
	w io.Writer
	f LogFormatter
}

func (self *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var log string

	start := time.Now()
	log = self.f.RequestLog(r, start)
	if len(log) > 0 {
		// TODO: Handle write error
		self.w.Write([]byte(log))
	}

	lw := &logResponseWriter{w, 0, 0}
	self.h.ServeHTTP(lw, r)

	log = self.f.ResponseLog(r, start, lw.status, lw.bytes)
	if len(log) > 0 {
		// TODO: Handle write error
		self.w.Write([]byte(log))
	}
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

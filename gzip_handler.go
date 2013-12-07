package httputil2

// TODO: Implement the http.Hijacker interface
// TODO: If a Content-Length header has been set, buffer the response until it's
//       all written and then flush ?

import (
	"compress/gzip"
	"net/http"
	"strings"
	"sync"
)

func GzipHandler(h http.Handler) http.Handler {
	return &gzipHandler{h}
}

type gzipHandler struct {
	h http.Handler
}

func (self *gzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Tell proxies that the content might vary
	w.Header().Add(HeaderVary, HeaderAcceptEncoding)

	// Ignore the client if it doesn't support the gzip content encoding
	if !strings.Contains(r.Header.Get(HeaderAcceptEncoding), "gzip") {
		self.h.ServeHTTP(w, r)
		return
	}

	f, f_ok := w.(http.Flusher)
	if !f_ok {
		panic("ResponseWriter is not a Flusher")
	}

	cn, cn_ok := w.(http.CloseNotifier)
	if !cn_ok {
		panic("ResponseWriter is not a CloseNotifier")
	}

	//
	w.Header().Set(HeaderContentEncoding, "gzip")
	gz := gzip.NewWriter(w)
	gzw := &gzipResponseWriter{
		gz: gz,
		w:  w,
		f:  f,
		cn: cn,
	}

	self.h.ServeHTTP(gzw, r)
	gzw.Flush()
}

type gzipResponseWriter struct {
	gz *gzip.Writer
	w  http.ResponseWriter
	f  http.Flusher
	cn http.CloseNotifier

	wroteHeader bool

	mu           sync.Mutex
	closeNotifyc chan bool
}

func (self *gzipResponseWriter) Header() http.Header {
	return self.w.Header()
}

func (self *gzipResponseWriter) WriteHeader(status int) {
	self.wroteHeader = true
	// Content-Length is wrong once compressed !
	self.Header().Del(HeaderContentLength)
	self.w.WriteHeader(status)
}

func (self *gzipResponseWriter) Write(p []byte) (int, error) {
	// We have to re-implement that logic that's in the http lib unfortunately
	if !self.wroteHeader {
		// Make sure to detect the content-type before we encode it
		// TODO: This should be done upstream instead
		if len(self.Header().Get(HeaderContentType)) == 0 {
			self.Header().Set(HeaderContentType, http.DetectContentType(p))
		}

		self.WriteHeader(http.StatusOK)
	}

	return self.gz.Write(p)
}

// For the http.Flusher interface. The server fails without.
func (self *gzipResponseWriter) Flush() {
	err := self.gz.Flush()
	// FIXME: How to deal with the error here ?
	if err != nil {
		panic(err)
	}
	self.f.Flush()
}

// For the the http.CloseNotifier interface. The server fails without.
func (self *gzipResponseWriter) CloseNotify() <-chan bool {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.closeNotifyc == nil {
		c := make(chan bool, 1)
		pc := self.cn.CloseNotify()

		go func() {
			<-pc
			// FIXME: Is this necessary ?
			self.gz.Close()
			c <- true
		}()
	}

	return self.closeNotifyc
}

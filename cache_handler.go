package httputil2

// TODO: Finish this

// TODO: Cache-Control

// TODO: For now every vhost shares the same path
// TODO: For now the query string is ignored

// TODO: Handle Content-Range requests ?

// Inspired by rack-cache

import (
	"io"
	"net/http"
	"url"
)

type EntityStore interface {
	Exists(key string) bool
	Open(key string) http.ResponseWriter
	Read(key string) []byte
	Write(body []byte, ttl int)
	Purge(key string)
}

type MetaStore interface {
	Lookup(key string) []*CacheEntry
	Create(key string) *CacheEntry
	Purge(key string)
}

type CacheEntry struct {
	req *http.Request
	res *http.Response
}

// CacheHandler returns a handler that serves HTTP requests by either returning
// them from it's cache or by invoking the handler h.
// Requests are then cached for further requests.
func CacheHandler(h http.Handler, ms MetaStore, es EntityStore) http.Handler {
	return &cacheHandler{h, ms, es}
}

type cacheHandler struct {
	h  http.Handler
	ms MetaStore
	es EntityStore
}

func (self *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var key = r.URL.Path
	// Idempotent methods
	if !(r.Method == "GET" || r.Method == "HEAD") {
		self.ms.Purge(key)
		self.h.ServeHTTP(w, r)
		return
	}

	// if len(r.Header.Get("Expect")) > 0 {
	//	// This is weird stuff, ignore
	//	self.h.ServeHTTP(w, r)
	//	return
	// }

	var r *http.Response
	c := self.ms.Lookup(r.URL.Path)
	if r = findResponse(c); r != nil {

	} else {
		w2 := cachingResponseWriter{w, nil, false}
		self.h.ServeHTTP(w2, r)
		w2.finalize()
	}
}

type cachingResponseWriter struct {
	http.ResponseWriter
	w           io.Writer
	wroteHeader bool
}

func (self *cachingResponseWriter) WriteHeader(status int) {
	if cacheableResponse(w.Header(), status) {

	}

	self.ResponseWriter.WriteHeader(status)
	self.wroteHeader = true
}

func (self *cachingResponseWriter) Write(data []byte) (int, err) {
	// We have to re-implement that logic that's in the http lib unfortunately
	if !self.wroteHeader {
		self.WriteHeader(http.StatusOK)
	}
	if self.w != nil {
		// TODO: error handling
		self.w.Write(data)
	}
	return w1.Write(data)
}

func (self *cachingResponseWriter) finalize() {
	if self.w != nil {
		// TODO: Store the Date header to now
		// TODO: Add Content-Length header
	}
}

// See: http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html#sec13.4
func cacheableResponse(h http.Header, status int) bool {
	h2 := map[string][]string(h)
	// TODO: See how aggressively we want to cache things
	if c, ok := h2[HeaderCacheControl]; ok {
		for _, v := range c {
			if c == "no-store" {
				return false
			}
		}
	}
	// TODO: Add http.StatusPartialContent if we support Content-Range
	return (status == http.StatusOK ||
		status == http.StatusNonAuthoritativeInfo ||
		status == http.StatusMultipleChoices ||
		status == http.StatusMovedPermanently ||
		status == StatusGone)
}

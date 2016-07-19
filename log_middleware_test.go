package httputil2

import (
	"net/http"
)

// Just there to verify at compile time that gzipResponseWriter implements the
// returning interfaces
func typecheck_logResponseWriter() (http.Flusher, http.CloseNotifier) {
	return new(logResponseWriter), new(logResponseWriter)
}

// The package sse implements the ServerSentEvent spec and provides a handler
// to write those type of events.
//
// Spec: https://html.spec.whatwg.org/multipage/comms.html#server-sent-events
//
// Is compatible with https://github.com/Yaffle/EventSource
package sse

import (
	"fmt"
	"net/http"
	"strings"
)

const LF = "\n"

// 2kb-sized comment
var IEPadding = []byte(fmt.Sprintf(":% 2048s\n", ""))

func IsIE(h http.Header) bool {
	// TODO: Make browser detection more accurate
	return strings.Contains(h.Get("User-Agent"), "MSIE ")
}

type EventWriter interface {
	http.ResponseWriter
	WriteEvent(e *Event) (int, error)
}

type Handler func(w EventWriter, r *http.Request, done <-chan bool)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	done := w.(http.CloseNotifier).CloseNotify()
	ew := newEW(w)

	if IsIE(r.Header) {
		w.Write(IEPadding) // 2kB padding for IE
		w.(http.Flusher).Flush()
	}
	// TODO: Make configurable?
	// io.WriteString(w, "retry: 2000\n")

	h(ew, r, done)
}

// Event writer
type ew struct {
	http.ResponseWriter
	flush func()
}

func newEW(w http.ResponseWriter) *ew {
	return &ew{w, w.(http.Flusher).Flush}
}

func (ew *ew) WriteEvent(e *Event) (int, error) {
	n, err := ew.ResponseWriter.Write([]byte(e.String()))
	ew.flush()
	return n, err
}

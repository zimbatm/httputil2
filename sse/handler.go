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

type Handler func(w EventWriter, r *http.Request, lastID string, closed <-chan bool)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	lastID := r.Header.Get("Last-Event-ID")
	if lastID == "" { // IE Fallback
		lastID = r.URL.Query().Get("lastEventId")
	}

	closed := w.(http.CloseNotifier).CloseNotify()
	ew := newEW(w)

	if IsIE(r.Header) {
		ew.Write(IEPadding) // 2kB padding for IE
	}
	// TODO: Make configurable?
	// io.WriteString(w, "retry: 2000\n")

	h(ew, r, lastID, closed)
}

// Event writer
type ew struct {
	http.ResponseWriter
	flush func()
}

func newEW(w http.ResponseWriter) *ew {
	return &ew{w, w.(http.Flusher).Flush}
}

func (ew *ew) Write(d []byte) (int, error) {
	n, err := ew.ResponseWriter.Write(d)
	ew.flush()
	return n, err
}

func (ew *ew) WriteEvent(e *Event) (int, error) {
	return ew.Write([]byte(e.String()))
}

package sse

// Represents a SSE event
//
// Because data has to be split by line we use an array, one element per line.
type Event struct {
	ID   string
	Name string
	Data []string // Each string is another line
}

// Builds an event where the data is a single-lined string
func MakeEvent(id, name, data string) *Event {
	return &Event{id, name, []string{data}}
}

// Returns the event in the SSE format
func (e *Event) String() (s string) {
	if len(e.ID) > 0 {
		s += "id: " + e.ID + LF
	}
	if len(e.Name) > 0 {
		s += "event: " + e.Name + LF
	}
	for _, data := range e.Data {
		s += "data: " + data + LF
	}
	return s + LF
}

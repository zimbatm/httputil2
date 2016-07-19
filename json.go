package httputil2

import (
	"encoding/json"
	"net/http"
)

// Data that's already been encoded to JSON
type JSONData []byte

// This is a NOOP
func (b JSONData) MarshalJSON() ([]byte, error) {
	return []byte(b), nil
}

const JSONContentType = "application/json"

func WriteJSON(w http.ResponseWriter, v interface{}) (int, error) {
	return WriteJSONWithStatus(w, StatusOK, v)
}

func WriteJSONWithStatus(w http.ResponseWriter, code int, v interface{}) (int, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}

	w.Header().Set(HeaderContentType, JSONContentType)
	w.WriteHeader(code)

	return w.Write(b)
}

// Standardize error messages
func WriteError(w http.ResponseWriter, code int, typ string, reason string) {
	WriteJSONWithStatus(w, code, map[string]string{
		"type":   typ,
		"reason": reason,
	})
}

func JSONNotFound(w http.ResponseWriter, r *http.Request) {
	WriteError(w, 404, "NotFound", r.URL.Path+" not found")
}

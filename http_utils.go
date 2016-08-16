package httputil2

import (
	"net/http"
)

// Checks each value of a header if it matches the value
func HeaderHas(h http.Header, key string, value string) bool {
	var vs []string
	var ok bool

	key = http.CanonicalHeaderKey(key)
	h2 := map[string][]string(h)
	if vs, ok = h2[key]; !ok {
		return false
	}
	for _, v := range vs {
		if v == value {
			return true
		}
	}
	return false
}

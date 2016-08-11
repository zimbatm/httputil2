package httputil2

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCleanPathMiddleware(t *testing.T) {
	var (
		err error
		r   *http.Request
		w   *httptest.ResponseRecorder
		h   http.Handler
	)

	h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	h = CleanPathMiddleware()(h)

	// Normal
	r, err = http.NewRequest(http.MethodGet, "/foo", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Error("Expected OK", w)
	}

	// Top-level
	r, err = http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Error("Expected OK", w)
	}

	// Trailing slash
	r, err = http.NewRequest(http.MethodGet, "/foo/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusTemporaryRedirect {
		t.Error("Expected redirect", w)
	}
	if l := w.Header().Get("Location"); l != "/foo" {
		t.Error("Expected /foo", l)
	}

	// Relative path
	r, err = http.NewRequest(http.MethodGet, "/../foo", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)

	if w.Code != http.StatusTemporaryRedirect {
		t.Error("Expected redirect", w)
	}
	if l := w.Header().Get("Location"); l != "/foo" {
		t.Error("Expected /foo", l)
	}
}

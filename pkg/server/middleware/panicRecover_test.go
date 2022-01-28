package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPanicRecoverer(t *testing.T) {
	var str bytes.Buffer

	log.SetOutput(&str)
	log.Print("test")

	mw := PanicRecoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("error")
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	mw.ServeHTTP(w, req)

	logged := str.String()
	if !strings.Contains(logged, "error") {
		t.Errorf("Output is invalid: %s", logged)
	}
}

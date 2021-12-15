package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestLogger(t *testing.T) {
	var str bytes.Buffer

	log.SetOutput(&str)
	log.Print("test")

	mw := RequestLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	mw.ServeHTTP(w, req)

	logged := str.String()
	if !strings.Contains(logged, "GET /") {
		t.Errorf("Output is invalid: %s", logged)
	}

}

package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestErrorLogger(t *testing.T) {
	var str bytes.Buffer

	log.SetOutput(&str)
	log.Print("test")

	mw := ErrorLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Foo error", http.StatusInternalServerError)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	mw.ServeHTTP(w, req)

	if http.StatusInternalServerError != w.Result().StatusCode {
		t.Error("Status missmatch")
	}

	logged := str.String()
	if !strings.Contains(logged, "Foo error\n") {
		t.Errorf("Output is invalid: %s", logged)
	}
}

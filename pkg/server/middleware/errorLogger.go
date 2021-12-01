package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func ErrorLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		buf := newLimitBuffer(512)
		ww.Tee(buf)

		defer func() {
			if ww.Status() < 400 {
				return
			}

			respBody, _ := ioutil.ReadAll(buf)
			log.Println(string(respBody))
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}

type limitBuffer struct {
	*bytes.Buffer
	limit int
}

func newLimitBuffer(size int) io.ReadWriter {
	return limitBuffer{
		Buffer: bytes.NewBuffer(make([]byte, 0, size)),
		limit:  size,
	}
}

func (b limitBuffer) Write(p []byte) (n int, err error) {
	if b.Buffer.Len() >= b.limit {
		return len(p), nil
	}
	limit := b.limit
	if len(p) < limit {
		limit = len(p)
	}
	return b.Buffer.Write(p[:limit])
}

func (b limitBuffer) Read(p []byte) (n int, err error) {
	return b.Buffer.Read(p)
}

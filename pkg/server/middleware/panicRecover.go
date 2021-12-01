package middleware

import (
	"log"
	"net/http"
)

func PanicRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recoveryMessage := recover(); recoveryMessage != nil {
				log.Println(recoveryMessage)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

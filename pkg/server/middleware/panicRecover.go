package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func PanicRecoverer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		defer func() {
			if recoveryMessage := recover(); recoveryMessage != nil {
				log.Println(recoveryMessage)
			}
		}()
		return next(context)
	}
}

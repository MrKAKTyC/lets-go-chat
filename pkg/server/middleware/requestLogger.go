package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		request := context.Request()
		log.Println(request.Method + " " + request.URL.String())
		return next(context)
	}
}

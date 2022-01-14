package serv

import (
	"github.com/MrKAKTyC/lets-go-chat/pkg/controller"
	"log"
	"net/http"

	serv "github.com/MrKAKTyC/lets-go-chat/pkg/generated"
	"github.com/MrKAKTyC/lets-go-chat/pkg/server/middleware"
	"github.com/labstack/echo/v4"
)

func Serve(port string, userController controller.User) {
	router := echo.New()

	router.Use(middleware.PanicRecoverer)
	router.Use(middleware.ErrorLogger)
	router.Use(middleware.RequestLogger)
	serv.RegisterHandlers(router, &userController)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Println(err)
	}
}

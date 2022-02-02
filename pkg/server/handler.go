package serv

import (
	"github.com/MrKAKTyC/lets-go-chat/pkg/controller"
	serv "github.com/MrKAKTyC/lets-go-chat/pkg/generated"
	"github.com/MrKAKTyC/lets-go-chat/pkg/server/middleware"
	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Serve(port string, userController controller.User) {
	router := echo.New()

	router.Use(middleware.PanicRecoverer)
	router.Use(middleware.ErrorLogger)
	router.Use(middleware.RequestLogger)
	echopprof.Wrap(router)
	serv.RegisterHandlers(router, &userController)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Println(err)
	}
}

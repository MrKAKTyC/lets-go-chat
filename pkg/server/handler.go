package serv

import (
	"fmt"
	"net/http"

	"github.com/MrKAKTyC/lets-go-chat/pkg/controller"
	serv "github.com/MrKAKTyC/lets-go-chat/pkg/generated"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"

	"github.com/labstack/echo/v4"
)

func Serve(port string) {
	fmt.Println("Running on port:", port)
	router := echo.New()
	server := &controller.User{Service: service.New()}

	serv.RegisterHandlers(router, server)

	http.ListenAndServe(":"+port, router)
}

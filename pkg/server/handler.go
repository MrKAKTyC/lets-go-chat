package serv

import (
	"fmt"
	"net/http"

	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
	userController "github.com/MrKAKTyC/lets-go-chat/pkg/server/controllers"
	"github.com/MrKAKTyC/lets-go-chat/pkg/services"

	"github.com/labstack/echo/v4"
)

func Serve(port string) {
	fmt.Println("Running on port:", port)
	router := echo.New()
	server := &userController.UserController{UserService: services.New(repository.UserRepositoryPGS())}

	RegisterHandlers(router, server)

	http.ListenAndServe(":"+port, router)
}

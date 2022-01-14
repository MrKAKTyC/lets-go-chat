package main

import (
	"github.com/MrKAKTyC/lets-go-chat/pkg/config"
	serv "github.com/MrKAKTyC/lets-go-chat/pkg/server"
)

func main() {
	conf := config.InitConfig()
	userController := InitializeController(conf)
	serv.Serve(conf.Server.Port, *userController)
}

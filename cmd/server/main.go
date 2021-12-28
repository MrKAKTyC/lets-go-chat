package main

import (
	"github.com/MrKAKTyC/lets-go-chat/pkg/config"
	serv "github.com/MrKAKTyC/lets-go-chat/pkg/server"
)

func main() {
	config := config.InitConfig()
	serv.Serve(config)
}

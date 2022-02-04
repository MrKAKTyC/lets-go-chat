package main

import (
	"bufio"
	"github.com/MrKAKTyC/lets-go-chat/pkg/config"
	serv "github.com/MrKAKTyC/lets-go-chat/pkg/server"
	"log"
	"os"
)

func main() {
	log.SetOutput(bufio.NewWriterSize(os.Stdout, 1024*16))
	conf := config.InitConfig()
	userController := InitializeController(conf)
	serv.Serve(conf.Server.Port, *userController)
}

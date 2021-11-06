package main

import (
	"os"

	serverPkg "github.com/MrKAKTyC/lets-go-chat/server/serv"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}
	serverPkg.Serve(port)
}

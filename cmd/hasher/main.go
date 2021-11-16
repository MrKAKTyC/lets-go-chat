package main

import (
	"fmt"

	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	hasher "github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"
)

func main() {
	repository := repository.UserPGS()
	service := service.New(repository)
	cur, err := service.Register(auth.CreateUserRequest{UserName: "user1", Password: "password"})
	if err == nil {
		fmt.Println(*cur.Id, *cur.UserName, err)
	}
	user, err := service.Authorize(auth.LoginUserRequest{UserName: "user2", Password: "password"})
	if err == nil {
		fmt.Println(user, err)
	}
}

func testHasher() {
	password := "securePassword"
	fmt.Println(hasher.HashPassword(password))
	hashedPassword, _ := hasher.HashPassword(password)
	fmt.Println(hashedPassword == "78da4a596a88bc5114f071ba590793bf3b37329d761230f33129983a747f414e") // wrong hash
	fmt.Println(hashedPassword == "debe062ddaaf9f8b06720167c7b65c778c934a89ca89329dcb82ca79d19e17d2") // carect hash
}

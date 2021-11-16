package main

import (
	"fmt"

	hasher "github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
)

func main() {
	testHasher()
}

func testHasher() {
	password := "securePassword"
	fmt.Println(hasher.HashPassword(password))
	hashedPassword, _ := hasher.HashPassword(password)
	fmt.Println(hashedPassword == "78da4a596a88bc5114f071ba590793bf3b37329d761230f33129983a747f414e") // wrong hash
	fmt.Println(hashedPassword == "debe062ddaaf9f8b06720167c7b65c778c934a89ca89329dcb82ca79d19e17d2") // carect hash
}

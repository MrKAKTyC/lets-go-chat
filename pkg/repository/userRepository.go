package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
)

type userRepository struct {
	connection pgx.Conn
}

func UserRepository() *userRepository {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:admin@127.0.0.1:5432/letsGoChat")
	if err != nil {
		log.Fatal(err)
	}
	return &userRepository{connection: *db}
}

func (repo userRepository) GetUser(login, password string) dao.User {
	var name string
	var weight string
	err := repo.connection.QueryRow(context.Background(), "SELECT * FROM users WHERE 'userName' LIKE '$1'", password).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
	return *(new(dao.User))
}

func (repo userRepository) CreateUser(login, password string) dao.User {

	return *(new(dao.User))
}

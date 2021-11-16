package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
)

type UserRepository interface {
	GetUser(login, password string) dao.User
	CreateUser(login, password string) dao.User
}

type userRepositoryPGS struct {
	connection pgx.Conn
}

func UserRepositoryPGS() *userRepositoryPGS {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:admin@127.0.0.1:5432/letsGoChat")
	if err != nil {
		log.Fatal(err)
	}
	return &userRepositoryPGS{connection: *db}
}

func (repo userRepositoryPGS) GetUser(login, password string) dao.User {
	user := dao.User{}
	err := repo.connection.QueryRow(context.Background(), "SELECT * FROM users WHERE 'userName' LIKE '$1'", password).Scan(user.Login, user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return *(new(dao.User))
}

func (repo userRepositoryPGS) CreateUser(login, password string) dao.User {
	userUUID := uuid.New().String()
	repo.connection.Exec(context.Background(), "INSERT INTO users (id, username, password) VALUES ('$1', '$2', '$3')", userUUID, login, password)

	return *(new(dao.User))
}

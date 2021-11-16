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

type User interface {
	Get(login, password string) (*dao.User, error)
	Create(login, password string) (*dao.User, error)
}

type userPGS struct {
	connection pgx.Conn
}

func UserPGS() *userPGS {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:admin@127.0.0.1:5432/letsGoChat")
	if err != nil {
		log.Fatal(err)
	}
	return &userPGS{connection: *db}
}

func (repo userPGS) Get(login, password string) (*dao.User, error) {
	user := dao.User{}
	err := repo.connection.QueryRow(context.Background(), "SELECT * FROM users WHERE username LIKE $1 AND password LIKE $2", login, password).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (repo userPGS) Create(login, password string) (*dao.User, error) {
	userUUID := uuid.New().String()

	_, err := repo.connection.Exec(context.Background(), "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", userUUID, login, password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	return &dao.User{ID: userUUID, Login: login}, nil
}

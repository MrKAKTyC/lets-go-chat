package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type userPGS struct {
	db sql.DB
}

func UserPGS(dbUrl string) *userPGS {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connection to", dbUrl, "established")
	}
	return &userPGS{db: *db}
}

func (r userPGS) Get(login, password string) (*dao.User, error) {
	user := dao.User{}
	err := r.db.QueryRow("SELECT * FROM users WHERE username LIKE $1 AND password LIKE $2", login, password).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (repo userPGS) Create(login, password string) (*dao.User, error) {
	userUUID := uuid.New().String()
	_, err := repo.db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", userUUID, login, password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	return &dao.User{ID: userUUID, Login: login}, nil
}

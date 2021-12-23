package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type userPGS struct {
	db *sql.DB
}

func UserPGS(db *sql.DB) *userPGS {
	return &userPGS{db: db}
}

func (repo userPGS) Get(login, password string) (*dao.User, error) {
	user := dao.User{}
	err := repo.db.QueryRow("SELECT id, username, password FROM users WHERE username LIKE $1 AND password LIKE $2", login, password).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		log.Printf("UserPGS::Get Query failed: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (repo userPGS) Create(login, password string) (*dao.User, error) {
	userUUID := uuid.New().String()
	_, err := repo.db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", userUUID, login, password)
	if err != nil {
		log.Printf("UserPGS::Create Query failed: %v\n", err)
		return nil, err
	}

	return &dao.User{ID: userUUID, Login: login}, nil
}

func (repo userPGS) GetLastOnline(userID string) (*time.Time, error) {
	lastOnline := new(time.Time)
	err := repo.db.QueryRow("SELECT lastonline FROM users WHERE id::text LIKE $1", userID).Scan(lastOnline)
	if err != nil {
		log.Printf("UserPGS::GetLastOnline Query failed: %v\n", err)
		return nil, err
	}
	return lastOnline, err
}

func (repo userPGS) UpdateLastOnline(userID string, logoutDate time.Time) error {
	_, err := repo.db.Exec("UPDATE users SET lastonline = $1 WHERE id::text LIKE $2", logoutDate, userID)
	return err
}

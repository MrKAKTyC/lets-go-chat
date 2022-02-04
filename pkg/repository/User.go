package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type UserPGS struct {
	db *sql.DB
}

type UserRepository interface {
	Get(login, password string) (*dao.User, error)
	Create(login, password string) (*dao.User, error)
	GetLastOnline(userID string) (*time.Time, error)
	UpdateLastOnline(userID string, logoutDate time.Time) error
}

func NewUserPGS(db *sql.DB) *UserRepository {
	var repo UserRepository
	repo = &UserPGS{db: db}
	return &repo
}

func (repo *UserPGS) Get(login, password string) (*dao.User, error) {
	user := dao.User{}
	err := repo.db.QueryRow("SELECT id, username, password FROM users WHERE username LIKE $1 AND password LIKE $2", login, password).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		log.Println("NewUserPGS::Get Query failed: ", err)
		return nil, err
	}

	return &user, nil
}

func (repo *UserPGS) Create(login, password string) (*dao.User, error) {
	userUUID := uuid.New().String()
	_, err := repo.db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", userUUID, login, password)
	if err != nil {
		log.Println("NewUserPGS::Create Query failed: ", err)
		return nil, err
	}

	return &dao.User{ID: userUUID, Login: login}, nil
}

func (repo *UserPGS) GetLastOnline(userID string) (*time.Time, error) {
	lastOnline := new(time.Time)
	err := repo.db.QueryRow("SELECT lastonline FROM users WHERE id::text LIKE $1", userID).Scan(lastOnline)
	if err != nil {
		log.Println("NewUserPGS::GetLastOnline Query failed: ", err)
		return nil, err
	}
	return lastOnline, err
}

func (repo *UserPGS) UpdateLastOnline(userID string, logoutDate time.Time) error {
	_, err := repo.db.Exec("UPDATE users SET lastonline = $1 WHERE id::text LIKE $2", logoutDate, userID)
	return err
}

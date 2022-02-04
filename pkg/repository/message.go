package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	_ "github.com/lib/pq"
)

type MessagePGS struct {
	db sql.DB
}

type MessageRepository interface {
	GetAfter(after time.Time) ([]*dao.Message, error)
	Create(message dao.Message) error
}

func NewMessagePGS(db *sql.DB) *MessageRepository {
	var repo MessageRepository
	repo = &MessagePGS{db: *db}
	return &repo
}

func (r *MessagePGS) GetAfter(after time.Time) ([]*dao.Message, error) {
	messages := make([]*dao.Message, 0)
	rows, err := r.db.Query("SELECT senderid, content, sendat FROM messages WHERE sendAt > $1", after)
	if err != nil {
		log.Println("[NewMessagePGS::GetAfter] Query failed: ", err)
		return nil, err
	}
	for rows.Next() {
		var sender, content string
		var date time.Time
		err = rows.Scan(&sender, &content, &date)
		if err != nil {
			log.Println(err)
		}
		message := &dao.Message{
			Sender:  sender,
			Content: content,
			Date:    date,
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (repo *MessagePGS) Create(message dao.Message) error {
	_, err := repo.db.Exec("INSERT INTO messages (senderID, content, sendat) VALUES ($1, $2, $3)", message.Sender, message.Content, message.Date)
	if err != nil {
		log.Println("[NewMessagePGS::Create] Query failed: ", err)
		return err
	}

	return nil
}

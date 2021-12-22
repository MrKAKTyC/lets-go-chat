package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	_ "github.com/lib/pq"
)

type messagePGS struct {
	db sql.DB
}

func MessagePGS(db *sql.DB) *messagePGS {
	return &messagePGS{db: *db}
}

func (r messagePGS) GetAfter(after time.Time) ([]*dao.Message, error) {
	messages := make([]*dao.Message, 0)
	rows, err := r.db.Query("SELECT senderid, content, sendat FROM messages WHERE sendAt > $1", after)
	if err != nil {
		log.Printf("[MessagePGS::GetAfter] Query failed: %v\n", err)
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

func (repo messagePGS) Create(message dao.Message) error {
	_, err := repo.db.Exec("INSERT INTO messages (senderID, content, sendat) VALUES ($1, $2, $3)", message.Sender, message.Content, message.Date)
	if err != nil {
		log.Printf("[MessagePGS::Create] Query failed: %v\n", err)
		return err
	}

	return nil
}

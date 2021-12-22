package dao

import "time"

type Message struct {
	Sender, Content string
	Date            time.Time
}

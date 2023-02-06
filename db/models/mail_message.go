package models

import (
	"time"
)

type MailMessage struct {
	Model
	ReadAt     *time.Time `json:"read_at,omitempty" gorm:"type:datetime"`
	Sender     string     `json:"sender" gorm:"type:varchar(255)"`
	Receiver   string     `json:"receiver" gorm:"type:varchar(255)"`
	RawMessage string     `json:"message"`
	MailBoxId  uint       `json:"mail_box_id"`
}

package models

import (
	"gorm.io/datatypes"
	"time"
)

type MailMessage struct {
	Model
	ReadAt     *time.Time     `json:"read_at,omitempty" gorm:"type:datetime"`
	Subject    string         `json:"subject"`
	Sender     string         `json:"sender" gorm:"type:varchar(255)"`
	Receiver   string         `json:"receiver" gorm:"type:varchar(255)"`
	RawMessage string         `json:"message"`
	MailBoxId  uint           `json:"mail_box_id"`
	Headers    datatypes.JSON `json:"headers"`
}

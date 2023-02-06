package models

import (
	"gorm.io/datatypes"
	"time"
)

type MailMessage struct {
	Model
	ReadAt     *time.Time     `json:"read_at,omitempty" gorm:"type:datetime"`
	Sender     string         `json:"sender"`
	Receiver   datatypes.JSON `json:"receiver"`
	RawMessage string         `json:"message"`
}

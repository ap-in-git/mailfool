package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time     `json:"created_at,omitempty" gorm:"type:datetime"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty" gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"  gorm:"type:datetime"`
}
type MailBox struct {
	Model
	Name        string `json:"name" gorm:"type:varchar(255)"`
	UserName    string `json:"user_name" gorm:"type:varchar(255)"`
	Password    string `json:"password" gorm:"type:varchar(255)"`
	TlsEnabled  bool   `json:"tls_enabled"`
	MaximumSize int8   `json:"maximum_size"`
}

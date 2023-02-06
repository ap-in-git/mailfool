package service

import (
	"github.com/ap-in-git/mailfool/db/models"
	"github.com/ap-in-git/mailfool/mailer/envelope"
	"gorm.io/gorm"
)

type MailMessageService struct {
	db *gorm.DB
}

func NewMailMessageService(db *gorm.DB) *MailMessageService {
	service := MailMessageService{db: db}
	return &service
}

func (s *MailMessageService) StoreEnvelope(envelope *envelope.Envelope) error {
	var mailMessage models.MailMessage
	mailMessage.MailBoxId = envelope.MailBox.ID
	mailMessage.Sender = envelope.Sender
	mailMessage.Receiver = envelope.Recipients[0]
	mailMessage.RawMessage = string(envelope.Data)
	s.db.Save(&mailMessage)
	return nil
}

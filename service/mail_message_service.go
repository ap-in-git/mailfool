package service

import (
	"github.com/ap-in-git/mailfool/db/models"
	"github.com/ap-in-git/mailfool/mailer/envelope"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"strings"
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
	rawMessage := string(envelope.Data)
	headers := getHeaders(rawMessage)
	bt, err := json.Marshal(&headers)
	if err == nil {
		mailMessage.Headers = bt
	}
	mailMessage.Subject = headers["Subject"]
	mailMessage.RawMessage = rawMessage
	s.db.Save(&mailMessage)
	return nil
}

func getHeaders(rawMessage string) map[string]string {
	headers := make(map[string]string, 0)
	for _, line := range strings.Split(strings.TrimRight(rawMessage, "\n"), "\n") {
		if line == "" {
			break
		}
		st := strings.Split(line, ":")
		if len(st) > 1 {
			headers[st[0]] = st[1]
		}
	}
	return headers
}

func (s *MailMessageService) GetAllMessageOfInbox(id int) []models.MailMessage {
	var messages []models.MailMessage
	s.db.Where("mail_box_id = ? ", id).Find(&messages)
	return messages
}

func (s *MailMessageService) FindMessage(id int) (*models.MailMessage, error) {
	var message models.MailMessage
	err := s.db.Take(&message, id).Error
	if err != nil {
		return nil, err
	}

	return &message, nil
}

package mailer

import (
	"github.com/ap-in-git/mailfool/db/models"
	"github.com/ap-in-git/mailfool/mailer/envelope"
)

type AuthHandler interface {
	IsValidLogin(authCredentials string) *models.MailBox
}

type EnvelopeHandler interface {
	StoreEnvelope(enveloper *envelope.Envelope) error
}

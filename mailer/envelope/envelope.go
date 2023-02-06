package envelope

import "github.com/ap-in-git/mailfool/db/models"

type Envelope struct {
	Sender     string
	Recipients []string
	Data       []byte
	MailBox    *models.MailBox
}

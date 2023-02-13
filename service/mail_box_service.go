package service

import (
	"fmt"
	"github.com/ap-in-git/mailfool/db/models"
	formRequest "github.com/ap-in-git/mailfool/form-request"
	"gorm.io/gorm"
)

type MailBoxService struct {
	db *gorm.DB
}

func NewMailBoxService(db *gorm.DB) *MailBoxService {
	service := MailBoxService{db: db}
	return &service
}

func (s *MailBoxService) CreateMailBox(request formRequest.MailBoxFormRequest) error {
	var mailBox models.MailBox
	mailBox.Name = request.Name
	mailBox.UserName = request.Username
	mailBox.Password = request.Password
	mailBox.TlsEnabled = request.TlsEnabled
	mailBox.MaximumSize = request.MaxSize
	err := s.db.Create(&mailBox).Error
	return err
}

func (s *MailBoxService) IsValidLogin(username string, password string) *models.MailBox {
	var mailBox models.MailBox
	err := s.db.Where("user_name = ? and password = ? ", username, password).Take(&mailBox).Error
	if err == nil {
		return &mailBox
	}
	if err != gorm.ErrRecordNotFound {
		fmt.Println(err.Error())
	}
	return nil
}

func (s *MailBoxService) GetAllMailBox() []models.MailBox {
	var boxes []models.MailBox
	s.db.Order("name asc").Order("created_at desc").Find(&boxes)
	return boxes
}

func (s *MailBoxService) CheckIfUsernameExist(userName string) bool {
	var mailbox models.MailBox
	s.db.Where("user_name = ?", userName).Take(&mailbox)
	return mailbox.ID > 0
}

func (s *MailBoxService) DeleteMailbox(id int) error {
	err := s.db.Delete(&models.MailBox{}, id).Error
	return err
}

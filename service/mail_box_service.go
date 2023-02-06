package service

import (
	"fmt"
	"github.com/ap-in-git/mailfool/db/models"
	"gorm.io/gorm"
	"strings"
)

type MailBoxService struct {
	db *gorm.DB
}

func NewMailBoxService(db *gorm.DB) *MailBoxService {
	service := MailBoxService{db: db}
	return &service
}

func (s *MailBoxService) CreateMailBox(mailBox models.MailBox) error {
	err := s.db.Create(&mailBox).Error
	return err
}

func (s *MailBoxService) IsValidLogin(authCredentials string) *models.MailBox {
	sp := strings.Split(authCredentials, ":")
	if len(sp) == 0 {
		return nil
	}
	var mailBox models.MailBox
	err := s.db.Where("user_name = ? and password = ? ", sp[0], sp[1]).Take(&mailBox).Error
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
	s.db.Order("name asc").Find(&boxes)
	return boxes
}

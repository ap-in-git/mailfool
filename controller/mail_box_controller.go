package controller

import (
	"github.com/ap-in-git/mailfool/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type MailBoxController struct {
	service *service.MailBoxService
}

func NewMailBoxController(db *gorm.DB) *MailBoxController {
	mailboxService := service.NewMailBoxService(db)
	mbc := MailBoxController{service: mailboxService}
	return &mbc
}

func (controller *MailBoxController) Index(c *gin.Context) {
	mailboxes := controller.service.GetAllMailBox()
	c.JSON(http.StatusOK, mailboxes)
}

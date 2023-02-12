package controller

import (
	"github.com/ap-in-git/mailfool/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type MailMessageController struct {
	service *service.MailMessageService
}

func (controller *MailMessageController) Index(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	messages := controller.service.GetAllMessageOfInbox(id)
	c.JSON(http.StatusOK, messages)
	return
}

func (controller *MailMessageController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mailMessage, err := controller.service.FindMessage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, mailMessage)
}

func NewMailMessageController(db *gorm.DB) *MailMessageController {
	s := service.NewMailMessageService(db)
	controller := MailMessageController{service: s}
	return &controller
}

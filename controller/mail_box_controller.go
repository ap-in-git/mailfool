package controller

import (
	formRequest "github.com/ap-in-git/mailfool/form-request"
	"github.com/ap-in-git/mailfool/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func (controller *MailBoxController) Store(c *gin.Context) {
	var mailBoxFormRequest formRequest.MailBoxFormRequest
	err := c.ShouldBindJSON(&mailBoxFormRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid form data",
		})
		return
	}
	userNameExist := controller.service.CheckIfUsernameExist(mailBoxFormRequest.Username)
	if userNameExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username already exist",
		})
		return
	}
	err = controller.service.CreateMailBox(mailBoxFormRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mailbox created successfully",
	})
}
func (controller *MailBoxController) Delete(c *gin.Context) {
	mailboxId, _ := strconv.Atoi(c.Param("id"))
	err := controller.service.DeleteMailbox(mailboxId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Mailbox deleted successfully",
	})
}

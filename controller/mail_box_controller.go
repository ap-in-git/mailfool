package controller

import (
	"github.com/ap-in-git/mailfool/service"
	"github.com/gin-gonic/gin"
)

type MailBoxController struct {
	service *service.MailBoxService
}

func (controller MailBoxController) Create(c *gin.Context) {

}

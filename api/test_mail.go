package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
)

func testMail(c *gin.Context) {
	println("Mail sent")
	from := "username:"
	password := "password"

	//toEmailAddress := "<paste the email address you want to send to>"
	to := []string{"ashokpoudel023@gmail.com"}

	host := "127.0.0.1"
	port := "2525"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, "ashok@bizzone.com", to, message)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "works",
	})
}

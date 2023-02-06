package api

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
	"net/http"
)

func testMail(c *gin.Context) {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "from@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "to@example.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("127.0.0.1", 2525, "username", ":password")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "works",
	})

	//from := "username:"
	//password := "password"

	////toEmailAddress := "<paste the email address you want to send to>"
	//to := []string{"bye@apinweb.com", "hi@apinweb.com"}
	//
	//host := "127.0.0.1"
	//port := "2525"
	//address := host + ":" + port
	//
	//subject := "Subject: This is the subject of the mail\n"
	//body := "This is the body of the mail"
	//message := []byte(subject + body)
	//
	//auth := smtp.PlainAuth("", from, password, host)
	//
	//err := smtp.SendMail(address, auth, "ashok@bizzone.com", to, message)
	//if err != nil {
	//	fmt.Errorf(err.Error())
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "works",
	//})
}

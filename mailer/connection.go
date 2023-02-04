package mailer

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

type Connection struct {
	reader      *bufio.Reader
	writer      *bufio.Writer
	scanner     *bufio.Scanner
	conn        net.Conn
	authService AuthService
	Envelope    *Envelope
}

func (c *Connection) writeSmtpMessage(statusCode int, message string) {
	_, err := c.writer.WriteString(strconv.Itoa(statusCode) + " " + message + "\r\n")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
func (c *Connection) writeWithDash(statusCode int, message string) {
	_, err := c.writer.WriteString(strconv.Itoa(statusCode) + "-" + message + "\r\n")
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (c *Connection) flushMessage() {
	err := c.writer.Flush()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func (c *Connection) reply(statusCode int, message string) {
	c.writeSmtpMessage(statusCode, message)
	c.flushMessage()
}

func (c *Connection) Serve() {
	c.sendWelcomeResponse()
}

func (c *Connection) sendWelcomeResponse() {
	c.reply(220, "SMTP Ready")

	for {
		for c.scanner.Scan() {
			c.handleResponse(c.scanner.Text())
		}
	}
}

func (c *Connection) handleResponse(line string) {
	sp := strings.Fields(line)

	//SMTP message should have at least two words.
	//One with smtp command and remaining being value needed to be processed
	if len(sp) < 2 {
		c.reply(502, "Command needs to be at least two words")
		return
	}
	switch sp[0] {
	case "EHLO":
		c.handleExtendedHello(sp)
	case "AUTH":
		c.handleAuth(sp)
	case "MAIL":
		c.handleMail(sp)
	case "RCPT":
		c.handleRCPT(sp)
	default:
		c.reply(502, "Invalid command")

	}
}

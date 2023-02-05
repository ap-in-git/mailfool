package mailer

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
	"strconv"
	"strings"
)

type Config struct {
	TLSConfig *tls.Config
}
type Connection struct {
	reader      *bufio.Reader
	writer      *bufio.Writer
	scanner     *bufio.Scanner
	conn        net.Conn
	authService AuthService
	Envelope    *Envelope
	config      Config
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
	switch sp[0] {
	case "EHLO":
		c.handleExtendedHello(sp)
	case "AUTH":
		c.handleAuth(sp)
	case "MAIL":
		c.handleMail(sp)
	case "RCPT":
		c.handleRCPT(sp)
	case "DATA":
		c.handleData(sp)
	default:
		c.reply(502, "Invalid command")

	}
}

func (c *Connection) handleData(sp []string) {
	if c.Envelope == nil || len(c.Envelope.Recipients) == 0 {
		c.reply(502, "Missing MAIL RCPT commands.")
		return
	}
	c.reply(354, "Go ahead. End your data with <CR><LF>.<CR><LF>")
}

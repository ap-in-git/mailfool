package mailer

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

type Connection struct {
	reader  *bufio.Reader
	writer  *bufio.Writer
	scanner *bufio.Scanner
	conn    net.Conn
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

func (c *Connection) Serve() {

}

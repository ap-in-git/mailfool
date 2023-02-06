package mailer

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/ap-in-git/mailfool/config"
	"github.com/ap-in-git/mailfool/db/models"
	"github.com/ap-in-git/mailfool/mailer/envelope"
	"io"
	"log"
	"net"
	"net/textproto"
	"strconv"
	"strings"
	"time"
)

type Connection struct {
	reader          *bufio.Reader
	writer          *bufio.Writer
	scanner         *bufio.Scanner
	conn            net.Conn
	authService     AuthHandler
	Envelope        *envelope.Envelope
	TLS             *tls.ConnectionState
	config          *config.MailConfig
	envelopeHandler EnvelopeHandler
	MailBox         *models.MailBox
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
	for {
		for c.scanner.Scan() {
			c.handleResponse(c.scanner.Text())
		}
	}
}

func (c *Connection) sendWelcomeResponse() {
	c.reply(220, "SMTP Ready")
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
	case "STARTTLS":
		c.handleSTARTTLS(sp)
	case "DATA":
		c.handleData(sp)
	case "QUIT":
		c.handleQUIT()

	default:
		c.reply(502, "Invalid command")

	}
}

func (c *Connection) handleData(sp []string) {
	if c.Envelope == nil || len(c.Envelope.Recipients) == 0 {
		c.reply(502, "Missing MAIL RCPT commands.")
		return
	}
	c.reply(354, "Go ahead. End your data with <CRLF>.<CRLF>")

	data := &bytes.Buffer{}
	reader := textproto.NewReader(c.reader).DotReader()

	_, err := io.CopyN(data, reader, c.config.MaxMessageSize)

	if err == io.EOF {
		c.Envelope.Data = data.Bytes()
		err := c.envelopeHandler.StoreEnvelope(c.Envelope)
		if err != nil {
			fmt.Println(err.Error())
			c.reply(502, "Something went wrong while storing data")
			return
		}
		c.reply(250, "Thank you.")
		c.reset()
		return
	}
	if err != nil {
		return
	}
}

func (c *Connection) handleSTARTTLS(sp []string) {

	if c.TLS != nil {
		c.reply(502, "Already running on TLS")
		return
	}

	if c.config.TLSConfig == nil {
		c.reply(502, "Tls not supported")
		return
	}
	c.reply(220, "Go ahead")
	tlsConn := tls.Server(c.conn, c.config.TLSConfig)

	// Perform a handshake
	if err := tlsConn.Handshake(); err != nil {
		c.reply(550, "Handshake error")
		return
	}

	// Reset envelope, new EHLO/HELO is required after STARTTLS
	c.reset()

	// Reset deadlines on the old connection - zero it out
	c.conn.SetDeadline(time.Time{})

	// Replace connection with a TLS connection
	c.conn = tlsConn
	c.reader = bufio.NewReader(c.conn)
	c.writer = bufio.NewWriter(c.conn)
	c.scanner = bufio.NewScanner(c.reader)

	state := tlsConn.ConnectionState()
	c.TLS = &state

	// Flush the connection to set up new timeout deadlines
	c.flush()
}

func (c *Connection) reset() {
	c.Envelope = nil
}

func (c *Connection) flush() {
	c.writer.Flush()
}

func (c *Connection) close() {
	c.writer.Flush()
	time.Sleep(200 * time.Millisecond)
	if c.conn != nil {
		c.conn.Close()
	}
}

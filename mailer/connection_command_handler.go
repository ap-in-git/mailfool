package mailer

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"strconv"
	"strings"
)

func (c *Connection) handleExtendedHello(sp []string) {
	size := 20 * 1024 * 1024 // 20 MB
	c.writeWithDash(250, sp[1])
	c.writeWithDash(250, "SIZE "+strconv.Itoa(size))

	if c.config.TLSConfig != nil && c.TLS == nil {
		c.reply(250, "STARTTLS")
		return
	}
	c.reply(250, "AUTH LOGIN PLAIN CRA-MD5")
	return
}

func (c *Connection) handleAuth(cmd []string) {
	if len(cmd) < 3 {
		c.reply(502, "Invalid number of parameter for auth")
		return
	}
	authDecoded, err := base64.StdEncoding.DecodeString(cmd[2])
	if err != nil {
		c.reply(502, "Illegal base64 data for auth credentials")
		return
	}
	authString := string(authDecoded)
	if !c.authService.IsValidLogin(authString) {
		c.reply(535, "Authentication failed")
		return
	}
	c.reply(235, "2.7.0 Authentication successful")
	return

}

func (c *Connection) handleMail(cmd []string) {
	params := strings.Split(cmd[1], ":")
	if len(params) < 2 {
		c.reply(502, "Invalid number of parameters")
		return
	}
	address, err := parseAddress(params[1], "sender")
	if err != nil {
		c.reply(502, err.Error())
		return
	}
	envelope := Envelope{
		Sender:     address,
		Recipients: []string{},
	}
	c.Envelope = &envelope
	c.reply(250, "Go ahead.")

}

func (c *Connection) handleRCPT(cmd []string) {
	if c.Envelope == nil {
		c.reply(502, "Missing MAIL FROM command.")
		return
	}

	params := strings.Split(cmd[1], ":")
	if len(params) < 2 {
		c.reply(502, "Invalid number of parameters")
		return
	}

	address, err := parseAddress(params[1], "receiver")
	if err != nil {
		c.reply(502, err.Error())
		return
	}
	c.Envelope.Recipients = append(c.Envelope.Recipients, address)
	c.reply(250, "Go ahead.")
}

func parseAddress(input string, inputType string) (string, error) {
	if len(input) < 3 {
		return "", fmt.Errorf("invalid "+inputType+" email length for %s", input)
	}
	if input[0] != '<' || input[len(input)-1] != '>' {
		return "", fmt.Errorf("invalid "+inputType+" email format. Should start with < and end with > for %s", input)
	}
	address, err := mail.ParseAddress(input[1 : len(input)-1])
	if err != nil {
		return "", err
	}
	//println(input)
	return address.Address, nil
}

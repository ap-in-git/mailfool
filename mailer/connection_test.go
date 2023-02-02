package mailer

import (
	"bufio"
	"bytes"
	"strconv"
	"testing"
)

func TestConnection_TestWriteSmtpMessage(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer: writer,
		reader: reader,
	}
	code := 220
	message := "smtp-ready"
	c.writeSmtpMessage(code, message)
	c.flushMessage()
	expectedMessage := strconv.Itoa(code) + " " + message + "\r\n"
	if b.String() != expectedMessage {
		t.Fatalf("Got %v Want %v", b.String(), expectedMessage)
	}
}

func TestConnection_TestWriteSmtpMessageWithDash(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer: writer,
		reader: reader,
	}
	code := 220
	message := "smtp-ready"
	c.writeWithDash(code, message)
	c.flushMessage()
	expectedMessage := strconv.Itoa(code) + "-" + message + "\r\n"
	if b.String() != expectedMessage {
		t.Fatalf("Got %v Want %v", b.String(), expectedMessage)
	}
}

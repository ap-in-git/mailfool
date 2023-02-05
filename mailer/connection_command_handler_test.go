package mailer

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"net"
	"testing"
)

func TestConnection_TestHandleExtendedHello(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer: writer,
		reader: reader,
	}
	host := "localhost"
	c.handleExtendedHello([]string{"EHLO", host})
	expected := "250-localhost\r\n250-SIZE 20971520\r\n250 AUTH LOGIN PLAIN CRA-MD5\r\n"
	if b.String() != expected {
		t.Fatalf("Got %v Want %v", b.String(), expected)
	}
}

type FakerService struct {
}

func (s FakerService) IsValidLogin(authDetails string) bool {
	return authDetails == "user:password"
}

func TestConnection_TestHandleAuth(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer:      writer,
		reader:      reader,
		authService: FakerService{},
	}
	username := "user"
	password := "password"
	authCredentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	username = "invalidUsername"
	password = "password"
	invalidAuthCredentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	testCases := []struct {
		commands         []string
		expectedResponse string
	}{
		{
			commands:         []string{"AUTH", "LOGIN", "Invalid base64 string"},
			expectedResponse: "502 Illegal base64 data for auth credentials\r\n",
		},
		{
			commands:         []string{"AUTH", "LOGIN", invalidAuthCredentials},
			expectedResponse: "535 Authentication failed\r\n",
		},
		{
			commands:         []string{"AUTH", "LOGIN", authCredentials},
			expectedResponse: "235 2.7.0 Authentication successful\r\n",
		},
	}

	for _, testCase := range testCases {
		c.handleAuth(testCase.commands)
		if testCase.expectedResponse != b.String() {
			t.Fatalf("Got %v Want %v", b.String(), testCase.expectedResponse)
		}
		b.Reset()
	}
}

func TestConnection_TestHandleMail(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer: writer,
		reader: reader,
	}
	testCases := []struct {
		commands         []string
		expectedResponse string
	}{
		{
			commands:         []string{"MAIL", "Invalid mail command"},
			expectedResponse: "502 Invalid number of parameters\r\n",
		},
		{
			commands:         []string{"MAIL", "FROM:34"},
			expectedResponse: "502 invalid sender email length for 34\r\n",
		},
		{
			commands:         []string{"MAIL", "FROM:hi@apinweb.com"},
			expectedResponse: "502 invalid sender email format. Should start with < and end with > for hi@apinweb.com\r\n",
		},
		{
			commands:         []string{"MAIL", "FROM:<hiapinwebcom>"},
			expectedResponse: "502 mail: missing '@' or angle-addr\r\n",
		},
		{
			commands:         []string{"MAIL", "FROM:<hi@apinwebcom>"},
			expectedResponse: "250 Go ahead.\r\n",
		},
	}

	for _, testCase := range testCases {
		c.handleMail(testCase.commands)
		if testCase.expectedResponse != b.String() {
			t.Fatalf("Got %v Want %v", b.String(), testCase.expectedResponse)
		}
		b.Reset()
	}
}

func TestConnection_TestRCPT(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer: writer,
		reader: reader,
	}

	envelope := Envelope{Sender: "hi@apinweb.com", Recipients: []string{}}
	testCases := []struct {
		commands         []string
		expectedResponse string
		envelope         *Envelope
	}{
		{
			commands:         []string{"RCPT", "Envelope not set"},
			expectedResponse: "502 Missing MAIL FROM command.\r\n",
			envelope:         nil,
		},
		{
			commands:         []string{"RCPT", "Invalid params"},
			expectedResponse: "502 Invalid number of parameters\r\n",
			envelope:         &envelope,
		},
		{
			commands:         []string{"RCPT", "TO:hi@apinweb.com"},
			expectedResponse: "502 invalid receiver email format. Should start with < and end with > for hi@apinweb.com\r\n",
			envelope:         &envelope,
		},
		{
			commands:         []string{"RCPT", "TO:<hiatinvalidemail.com>"},
			expectedResponse: "502 mail: missing '@' or angle-addr\r\n",
			envelope:         &envelope,
		},
		{
			commands:         []string{"RCPT", "TO:<hi@apinweb.com>"},
			expectedResponse: "250 Go ahead.\r\n",
			envelope:         &envelope,
		},
	}

	for _, testCase := range testCases {
		c.Envelope = testCase.envelope
		c.handleRCPT(testCase.commands)
		if testCase.expectedResponse != b.String() {
			t.Fatalf("Got %v Want %v", b.String(), testCase.expectedResponse)
		}
		b.Reset()
	}
}
func TestConnection_TestHandleQuit(t *testing.T) {
	_, client := net.Pipe()
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer: writer,
		reader: reader,
		conn:   client,
	}
	c.handleQUIT()
	expectedString := "221 OK, bye\r\n"
	if b.String() != expectedString {
		t.Fatalf("Got %v want %v", b.String(), expectedString)
	}

}

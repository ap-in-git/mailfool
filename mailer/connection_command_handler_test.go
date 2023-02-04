package mailer

import (
	"bufio"
	"bytes"
	"encoding/base64"
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

	expected := "250-localhost\r\n250-SIZE 20971520\r\n250 AUTH LOGIN\r\n"
	if b.String() != expected {
		t.Fatalf("Got %v Want %v", b.String(), expected)
	}
}
func TestConnection_HandleInvalidCommand(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer: writer,
		reader: reader,
	}

	c.handleResponse("INVALID")

	expected := "502 Command needs to be at least two words\r\n"
	if b.String() != expected {
		t.Fatalf("Got %v Want %v", b.String(), expected)
	}
	b.Reset()
	c.handleResponse("INVALID OTHER")
	expected = "502 Invalid command\r\n"
	if b.String() != expected {
		t.Fatalf("Got %v Want %v", b.String(), expected)
	}
}

type TestAuthService struct {
}

func (s TestAuthService) IsValidLogin(authDetails string) bool {
	return authDetails == "user:password"
}

func TestConnection_TestHandleAuth(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	reader := bufio.NewReader(&b)
	c := Connection{
		writer:      writer,
		reader:      reader,
		authService: TestAuthService{},
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
			commands:         []string{"AUTH", "TEST"},
			expectedResponse: "502 Invalid number of parameter for auth\r\n",
		},
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
			commands:         []string{"RCPT", "FROM:34"},
			expectedResponse: "250 Go ahead\r\n",
			envelope:         &envelope,
		},
		//{
		//	commands:         []string{"RCPT", "FROM:hi@apinweb.com"},
		//	expectedResponse: "502 invalid sender email format. Should start with < and end with > for hi@apinweb.com\r\n",
		//	envelope:         &envelope,
		//},
		//{
		//	commands:         []string{"RCPT", "FROM:<hiapinwebcom>"},
		//	expectedResponse: "502 mail: missing '@' or angle-addr\r\n",
		//	envelope:         &envelope,
		//},
		//{
		//	commands:         []string{"RCPT", "FROM:<hi@apinwebcom>"},
		//	expectedResponse: "250 Go ahead.\r\n",
		//	envelope:         &envelope,
		//},
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
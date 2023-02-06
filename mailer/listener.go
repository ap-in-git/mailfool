package mailer

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/ap-in-git/mailfool/config"
	"github.com/ap-in-git/mailfool/service"
	"log"
	"net"
	"path/filepath"
	"runtime"
)

func ListenMailConnection(mailConfig config.MailConfig, service *service.MailBoxService, messageService *service.MailMessageService) {
	host := mailConfig.Host
	port := mailConfig.Port
	networkUrl := host + ":" + port
	ln, err := net.Listen("tcp", networkUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Listening on host %s and port %s\n", host, port)
	host, port, err = net.SplitHostPort(ln.Addr().String())
	if mailConfig.Tls {
		cer, err := generateTlsCertificate()
		if err != nil {
			log.Fatal(err)
			return
		}
		tlsConfig := tls.Config{Certificates: []tls.Certificate{cer}}
		mailConfig.TLSConfig = &tlsConfig
	}
	for {
		acceptIncomingConnection(ln, mailConfig, service, messageService)
	}

}

type TempAuthService struct {
}

func (s TempAuthService) IsValidLogin(authCredentials string) bool {
	return true
}

func acceptIncomingConnection(ln net.Listener, mailConfig config.MailConfig, service *service.MailBoxService, messageService *service.MailMessageService) {
	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf(err.Error())
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(reader)
	if err != nil {
		log.Println(err.Error())
		return
	}

	sc := &Connection{
		conn:            conn,
		reader:          reader,
		writer:          writer,
		scanner:         scanner,
		authService:     service,
		envelopeHandler: messageService,
		config:          &mailConfig,
	}

	sc.config = &mailConfig
	go sc.Serve()

}

func generateTlsCertificate() (tls.Certificate, error) {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	certKeyFileLocation := basePath + "/../certs/server.rsa.crt"
	certKeyLocation := basePath + "/../certs/server.rsa.key"
	cer, err := tls.LoadX509KeyPair(certKeyFileLocation, certKeyLocation)
	return cer, err
}

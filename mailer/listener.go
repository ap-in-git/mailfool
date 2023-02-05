package mailer

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"runtime"
)

func ListenMailConnection() {
	host := "127.0.0.1"
	port := "2525"
	networkUrl := host + ":" + port
	ln, err := net.Listen("tcp", networkUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Listening on host %s and port %s\n", host, port)
	host, port, err = net.SplitHostPort(ln.Addr().String())
	defer ln.Close()

	for {
		acceptIncomingConnection(ln)
		//conn, err := ln.Accept()
	}

}

type TempAuthService struct {
}

func (s TempAuthService) IsValidLogin(authCredentials string) bool {
	return true
}

func acceptIncomingConnection(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf(err.Error())
	}

	cer, err := generateTlsCertificate()
	if err != nil {
		log.Fatal(err)
		return
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(reader)
	if err != nil {
		log.Println(err)
		return
	}

	cfg := tls.Config{
		Certificates: []tls.Certificate{cer},
	}
	sc := &Connection{
		conn:        conn,
		reader:      reader,
		writer:      writer,
		scanner:     scanner,
		authService: TempAuthService{},
		config: Config{
			TLSConfig: &cfg,
		},
	}
	sc.Serve()

}

func generateTlsCertificate() (tls.Certificate, error) {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	certKeyFileLocation := basePath + "/../certs/server.rsa.crt"
	certKeyLocation := basePath + "/../certs/server.rsa.key"
	cer, err := tls.LoadX509KeyPair(certKeyFileLocation, certKeyLocation)
	return cer, err
}

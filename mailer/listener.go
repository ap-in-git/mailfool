package mailer

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(reader)
	sc := &Connection{
		conn:        conn,
		reader:      reader,
		writer:      writer,
		scanner:     scanner,
		authService: TempAuthService{},
		config:      Config{},
	}
	sc.Serve()

}

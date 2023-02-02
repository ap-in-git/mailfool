package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/ap-in-git/mailfool/api"
	"github.com/joho/godotenv"
	"net"
	"strconv"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env files")
	}
	go api.InitializeApiRoutes()
	listenMail()

}

func listenMail() {
	ln, err := net.Listen("tcp", "127.0.0.1:2525")
	if err != nil {
		fmt.Println("The following error occured", err)
	} else {
		fmt.Println("The listener object has been created:", ln)
	}
	defer ln.Close()

	host, port, err := net.SplitHostPort(ln.Addr().String())
	fmt.Printf("Listening on host: %s, port: %s\n", host, port)

	if err != nil {
		panic(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)
		scanner := bufio.NewScanner(reader)
		writeMessage(writer, 220, "SMTP ready")
		writer.Flush()

		for {
			for scanner.Scan() {
				println(scanner.Text() + "===><ASd")
				handleScan(scanner.Text(), writer)
				writer.Flush()
				// Handle the textual version of the message
				//c.handle(c.scanner.Text())
			}
			err = scanner.Err()
			if err != nil && err == bufio.ErrTooLong {
				writeMessage(writer, 500, "Line too long")
			}
		}

		// Listen for an incoming connection

		//conn.re

	}

}

func handleScan(t string, writer *bufio.Writer) {
	println(t + "=======>")
	fields := strings.Fields(t)
	switch fields[0] {
	case "HELO":
		writeMessage(writer, 250, "Everything went well")
		//writer.Flush()
	case "EHLO":
		writer.WriteString("250-" + "localhost" + "\r\n")
		writer.WriteString("250-" + "SIZE 1000000" + "\r\n")
		//writer.WriteString("250-" + "AUTH LOGIN" + "\r\n")
		//writer.WriteString("250 " + "AUTH" + "\r\n")
		//writer.WriteString("250-" + "8BITMIME" + "\r\n")
		writeMessage(writer, 250, "AUTH LOGIN")
	case "AUTH":
		sDec, _ := base64.StdEncoding.DecodeString(fields[2])
		//writer.WriteString("250-" + "AUTH LOGIN" + "\r\n")
		//writer.WriteString("250-" + "AUTH LOGIN" + "\r\n")

		writeMessage(writer, 235, "2.7.0 Authentication successful")
		//check username password here
		println(string(sDec) + "xxxxx")
	}
}

func writeMessage(writer *bufio.Writer, code int, message string) {
	writer.WriteString(strconv.Itoa(code) + " " + message + "\r\n")
}

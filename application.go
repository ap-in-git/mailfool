package main

import (
	"github.com/ap-in-git/mailfool/api"
	"github.com/ap-in-git/mailfool/mailer"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env files")
	}
	go api.InitializeApiRoutes()
	mailer.ListenMailConnection()
	//http.StatusOK
	//listenMail()

}

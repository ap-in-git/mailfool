package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ap-in-git/mailfool/api"
	"github.com/ap-in-git/mailfool/config"
	"github.com/ap-in-git/mailfool/connection"
	"github.com/ap-in-git/mailfool/db/models"
	"github.com/ap-in-git/mailfool/mailer"
	"github.com/ap-in-git/mailfool/service"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	appConfig := getAppConfig()
	go api.InitializeApiRoutes()
	database, gormDb := connection.Init(&appConfig.Db)
	defer database.Close()
	//setupDefaultValues(gormDb)
	mailBoxService := service.NewMailBoxService(gormDb)
	mailMessageService := service.NewMailMessageService(gormDb)
	mailer.ListenMailConnection(appConfig.Mail, mailBoxService, mailMessageService)
}

func setupDefaultValues(db *gorm.DB) {
	t := models.MailBox{
		Name:        "Test",
		UserName:    "username",
		Password:    "password",
		TlsEnabled:  true,
		MaximumSize: 10,
	}
	db.Create(&t)

}

func getAppConfig() config.AppConfig {
	f := "config.toml"
	if _, err := os.Stat(f); err != nil {
		log.Fatalf(err.Error())
	}
	var appConfig config.AppConfig
	_, err := toml.DecodeFile(f, &appConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return appConfig
}

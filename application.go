package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ap-in-git/mailfool/api"
	"github.com/ap-in-git/mailfool/config"
	"github.com/ap-in-git/mailfool/connection"
	"github.com/ap-in-git/mailfool/db/models"
	"github.com/ap-in-git/mailfool/mailer"
	"github.com/ap-in-git/mailfool/service"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	appConfig := getAppConfig()
	sqlDb, db := connection.Init(&appConfig.Db)
	defer sqlDb.Close()
	//setupDefaultddValues(db)
	go api.InitializeApiRoutes(db)
	mailBoxService := service.NewMailBoxService(db)
	mailMessageService := service.NewMailMessageService(db)
	mailer.ListenMailConnection(appConfig.Mail, mailBoxService, mailMessageService)
}

func setupDefaultValues(db *gorm.DB) {

	fk := faker.New()
	for i := 0; i < 12; i++ {
		mbox := models.MailBox{
			Name:        fk.Person().Name(),
			UserName:    fk.Internet().User(),
			Password:    "password",
			MaximumSize: int8(i + 1),
			TlsEnabled:  fk.Bool(),
		}
		db.Create(&mbox)
	}
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

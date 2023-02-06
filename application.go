package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ap-in-git/mailfool/api"
	"github.com/ap-in-git/mailfool/config"
	"github.com/ap-in-git/mailfool/connection"
	"github.com/ap-in-git/mailfool/mailer"
	"log"
	"os"
)

func main() {
	appConfig := getAppConfig()
	go api.InitializeApiRoutes()
	database, _ := connection.Init(&appConfig.Db)
	defer database.Close()
	mailer.ListenMailConnection(appConfig.Mail)
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

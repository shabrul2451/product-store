package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DatabaseConnectionString string
	DbUsername               string
	DbPassword               string
	DbServer                 string
	DbPort                   string
	DatabaseName             string
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
		return
	}
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbServer = os.Getenv("DB_SERVER")
	DbPort = os.Getenv("DB_PORT")
	DatabaseName = os.Getenv("DB_NAME")
	if DbUsername != "" {
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort + "/?authSource=admin&readPreference=primary&directConnection=true&ssl=false"
	} else {
		DatabaseConnectionString = "mongodb://localhost:27017"
	}
}

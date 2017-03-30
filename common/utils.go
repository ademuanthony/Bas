package common

import (
	"os"
	"log"
	"encoding/json"
)

type configuration struct {
	DbHost, DbUserName, DbPassword, Database, Server, Port string
}

//Initialize AppConfig
var AppConfig configuration

func initConfig() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil{
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil{
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}
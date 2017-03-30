package common

import (
	"os"
	"log"
	"encoding/json"
	"net/http"
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


type (
	appError struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}
	errorResponse struct {
		Data appError `json:"data"`
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application.json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResponse{Data: errObj}); err == nil{
		w.Write(j)
	}
}
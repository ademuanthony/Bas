package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"reflect"
	"errors"
	"github.com/ademuanthony/bas/resources"
)

type configuration struct {
	DbHost, DbUserName, DbPassword, Database, Server, Port string
	TokenLifeTime int
}

//Initialize AppConfig
var AppConfig configuration

func initConfig() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResponse struct {
		Data appError `json:"data"`
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application.json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResponse{Data: errObj}); err == nil {
		w.Write(j)
	}
}

func SendResult(w http.ResponseWriter, response resources.ResponseResource, statusCode int) {
	if j, err := json.Marshal(response); err != nil {
		DisplayAppError(w, err, "An unexpected error has eccured", 500)
	} else {
		response.StatusCode = statusCode
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(j)
	}
}

func Log(key string, data interface{})  {
	fmt.Printf("%s: %v\n", key, data)
}


func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

func FillStruct(s interface{}, m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}



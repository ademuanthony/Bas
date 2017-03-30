package common

import (
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
)

// using asymmetric crypto/RSA keys
const (
	// openssl genrsa -out app.rsa
	privatePath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout app.rsa.pub
	publicPath = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// Read key files before starting hppt handlers
func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privatePath)
	if err != nil{
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(publicPath)
	if err != nil{
		log.Fatalf("[initKeys}: %s\n", err)
	}
}

// Generate JWT token
func GenerateJWT(name, role string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	//set claim for JWT token
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// set the expire time for JWT token
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil{
		return "", err
	}
	return tokenString, nil
}


func Authorize(w http.ResponseWriter, r *http.Request) {
	// validate the token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		//Verufy the token with public key, which is the counter part of private key
		return verifyKey, nil
	})

	if err != nil{
		switch err.(type) {
			case *jwt.ValidationError: // JWT Validation error


			return

		default:

			return

		}
	}
}

package common

import (
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
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

	claims := make(jwt.MapClaims)

	//set claim for JWT token
	claims["iss"] = "admin"
	claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// set the expire time for JWT token
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	t.Claims = claims

	tokenString, err := t.SignedString(signKey)
	if err != nil{
		return "", err
	}
	return tokenString, nil
}


func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// validate the token
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err != nil{
		switch err.(type) {
			case *jwt.ValidationError: // JWT Validation error
				validationError := err.(*jwt.ValidationError)

				switch validationError.Errors {
					case jwt.ValidationErrorExpired: //JWT expired
						DisplayAppError(w, err, "Access Token is expired, Get a new one", 401)
						return
					default:
						DisplayAppError(w, err, "Error while parsing Access Token", 500)
						return
				}
		default:
			DisplayAppError(w, err, "Error while parsing Access Token", 500)
			return

		}
	}
	if token.Valid{
		next(w, r)
	}else{
		DisplayAppError(w, err, "Invalid Access Token", 401)
		return
	}
}

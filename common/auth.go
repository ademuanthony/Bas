package common

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"context"
	"time"
)

// using asymmetric crypto/RSA keys
const (
	// openssl genrsa -out app.rsa
	privatePath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	publicPath = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	//verifyKey, signKey []byte

	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type TokenData struct {
	UserId int64
	Permissions []int64
}


// Read key files before starting http handlers
func initKeys() {
	var err error

	signBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

// Generate JWT token
func GenerateJWT(tokenData TokenData) (string, error) {

	claims := make(jwt.MapClaims)

	//set claim for JWT token
	claims["iss"] = "admin"
	claims["UserInfo"] = tokenData

	// set the expire time for JWT token
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(AppConfig.TokenLifeTime)).Unix()
	claims["iat"] = time.Now().Unix()

	//t.Claims = claims

	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := t.SignedString(signKey)
	if err != nil {
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

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT Validation error
			validationError := err.(*jwt.ValidationError)

			switch validationError.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(w, err, "Access Token is expired, Get a new one", 401)
				return
			default:
				DisplayAppError(w, err, "Error while parsing Access Token", http.StatusBadRequest)
				return
			}
		default:
			DisplayAppError(w, err, "Error while parsing Access Token", 500)
			return

		}
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		r = r.WithContext(context.WithValue(r.Context(), "UserInfo", claims["UserInfo"]))
		next(w, r)
	} else {
		DisplayAppError(w, err, "Invalid Access Token", 401)
		return
	}
}

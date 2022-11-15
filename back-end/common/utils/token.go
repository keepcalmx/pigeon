package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/yaml.v3"
)

var MySecret []byte

type MyClaims struct {
	UUID string `json:"user_id"`
	jwt.RegisteredClaims
}

func init() {
	yf, err := os.ReadFile("./config/jwt.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config map[any]any
	err = yaml.Unmarshal(yf, &config)
	if err != nil {
		log.Fatal(err)
	}
	MySecret = []byte(config["secret"].(string))
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	}
}

func GenToken(uuid string) (string, error) {
	claims := MyClaims{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(MySecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseToken(signedToken string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("this input param is not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("this token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("this token is not valid yet")
			} else {
				return nil, errors.New("otehr error with this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("cound not handle this token")
}

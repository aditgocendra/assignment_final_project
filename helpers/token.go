package helpers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = "UH!@OIH!@KJWD(!@(OPD)!_+!@(*#!@?>?~*!(&*DS4D"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id" : id,
		"email" : email,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	fmt.Println(headerToken)
	bearer := strings.HasPrefix(headerToken, "Bearer")

	fmt.Println(bearer)

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secret), nil
	})

	return token.Claims.(jwt.MapClaims), nil
}
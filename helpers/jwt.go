package helpers

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const key = "S3CR3T"

func GenerateToken(id, role uint, email string) (res string) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := parseToken.SignedString([]byte(key))
	if err != nil {
		log.Println("error parse token")
		return
	}
	return
}

func VerifyToken(ctx *gin.Context) (res interface{}, err error) {
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		err = errors.New("invalid header authorization")
		return
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (res interface{}, err error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.New("invalid method header alg")
			return
		}
		res, err = []byte(key), nil
		return
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		err = errors.New("token invalid")
		return
	}

	res = token.Claims.(jwt.MapClaims)
	return

}

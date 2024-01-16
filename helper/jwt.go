package helper

import (
	"ej-ex3-backend/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(user model.User) string {

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, _ := token.SignedString([]byte("ACCESS_SECRET_KEY"))

	return accessToken

}

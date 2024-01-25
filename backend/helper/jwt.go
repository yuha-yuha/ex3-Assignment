package helper

import (
	"ej-ex3-backend/model"
	"fmt"
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

func CheckJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte("ACCESS_SECRET_KEY"), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

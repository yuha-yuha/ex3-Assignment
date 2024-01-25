package database

import (
	"ej-ex3-backend/helper"
	"ej-ex3-backend/model"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/logger"
)

func GetUser(email string) model.User {
	db := DBconn()

	var user model.User

	db.Where("email = ?", email).First(&user)

	return user
}

func CreateUser(user model.User) error {

	var err error

	db := DBconn()

	BytesPassword := []byte(user.Password)
	HashedPassword, _ := bcrypt.GenerateFromPassword(BytesPassword, 5)
	user.Password = string(HashedPassword)

	err = db.Create(&user).Error

	return err

}

func FindUserById(id int64) (model.User, error) {
	var err error
	db := DBconn()
	var user model.User

	db.Logger = db.Logger.LogMode(logger.Info)

	err = db.Preload("UserWords").First(&user, id).Error

	return user, err
}

func CurrentUser(ctx *gin.Context, user *model.User) {

	tokenStr, err := ctx.Cookie("token")
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(403)
		return
	}

	token, err := helper.CheckJWT(tokenStr)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(403)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		*user, err = FindUserById(int64(claims["sub"].(float64)))

		if err != nil {
			ctx.AbortWithError(403, err)
			return
		}

	} else {
		ctx.AbortWithStatus(403)
	}

}

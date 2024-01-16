package handler

import (
	"ej-ex3-backend/helper"
	"ej-ex3-backend/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {

	var RequestJson struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.BindJSON(&RequestJson)
	if err != nil {
		ctx.Status(403)
		ctx.Abort()
	}

	user := model.GetUser(RequestJson.Email)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(RequestJson.Password))
	if err != nil {
		ctx.Status(403)
		ctx.Abort()
	}

	token := helper.CreateJWT(user)

	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)

}

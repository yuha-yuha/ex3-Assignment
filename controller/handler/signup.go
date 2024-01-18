package handler

import (
	"ej-ex3-backend/database"
	"ej-ex3-backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var SignUpRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := ctx.ShouldBindJSON(&SignUpRequest)

	if err != nil {
		log.Println(err)
		ctx.Abort()
	}

	log.Println(SignUpRequest)

	user := model.User{
		Name:     SignUpRequest.Name,
		Email:    SignUpRequest.Email,
		Password: SignUpRequest.Password,
	}

	if err = database.CreateUser(user); err != nil {
		log.Println(err)
		ctx.Status(400)
	}

	ctx.Status(200)

}

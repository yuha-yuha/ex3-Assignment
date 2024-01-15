package handler

import (
	"github.com/gin-gonic/gin"
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

}

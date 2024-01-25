package middleware

import (
	"ej-ex3-backend/helper"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthCheckMiddleware(ctx *gin.Context) {
	jwtStr, err := ctx.Cookie("token")

	if err != nil {
		log.Println(err)
		ctx.Status(403)
		ctx.Abort()
		return
	}

	_, err = helper.CheckJWT(jwtStr)

	if err != nil {
		log.Println(err)
		ctx.Status(403)
		ctx.Abort()
		return
	}

	ctx.Next()

}

package middleware

import (
	"ej-ex3-backend/helper"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthCheckMiddleware(ctx *gin.Context) {
	//jwtStr, err := ctx.Cookie("token")
	//フロントと時間の関係でCookieの使用はやめてURLクエリを使用
	token := ctx.Query("token")

	_, err := helper.CheckJWT(token)

	if err != nil {
		log.Println("aiueo", err)
		ctx.Status(403)
		ctx.Abort()
		return
	}

	ctx.Next()

}

package controller

import (
	"ej-ex3-backend/controller/handler"
	"ej-ex3-backend/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/login", handler.Login)
	r.POST("/signup", handler.SignUp)

	apiRouter := r.Group("/api/")
	{
		authRouter := apiRouter.Group("/auth")
		authRouter.Use(middleware.AuthCheckMiddleware)
		{
			authRouter.GET("/hello", func(ctx *gin.Context) {
				ctx.String(200, "hello")
			})
		}
	}
	return r
}

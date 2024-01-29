package controller

import (
	"ej-ex3-backend/controller/handler"
	"ej-ex3-backend/database"
	"ej-ex3-backend/middleware"
	"ej-ex3-backend/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
		},
		AllowCredentials: true,

		AllowHeaders: []string{"content-type"},
	}))

	apiRouter := r.Group("/api/")
	{
		apiRouter.POST("/login", handler.Login)
		apiRouter.POST("/signup", handler.SignUp)
		apiRouter.GET("/users/:id", handler.GetUserById)
		apiRouter.GET("/users/:id/words", handler.GetALLWordsByUser)
		apiRouter.GET("/words", handler.GetALLWordsByUsers)
		authRouter := apiRouter.Group("/auth")
		authRouter.Use(middleware.AuthCheckMiddleware)
		{
			authRouter.GET("/quizes", func(ctx *gin.Context) {
				var user model.User
				database.CurrentUser(ctx, &user)

				ctx.JSON(200, user)
			})
			authRouter.POST("/words", handler.CreateWordByUser)
		}
	}

	return r
}

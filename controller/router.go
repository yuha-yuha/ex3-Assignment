package controller

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/auth")
	return r
}

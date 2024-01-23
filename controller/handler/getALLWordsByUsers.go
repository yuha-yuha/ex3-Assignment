package handler

import (
	"ej-ex3-backend/database"
	"log"

	"github.com/gin-gonic/gin"
)

func GetALLWordsByUsers(ctx *gin.Context) {
	uw, err := database.GetALLUserWords()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(500)
	}

	ctx.JSON(200, uw)
}

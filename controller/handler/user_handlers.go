package handler

import (
	"ej-ex3-backend/database"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Println(id)
		ctx.AbortWithStatus(400)
		log.Println("errrrrrr")
		return
	}

	log.Println("aiueo")
	user, err := database.FindUserById(int64(id))

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(200, user)

}

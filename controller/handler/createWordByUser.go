package handler

import (
	"ej-ex3-backend/database"
	"ej-ex3-backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateWordByUser(ctx *gin.Context) {
	var word model.UserWord
	var user model.User
	var RequestJson struct {
		Japanese string `json:"japanese"`
		English  string `json:"english"`
	}

	err := ctx.ShouldBindJSON(&RequestJson)

	word.Japanese = RequestJson.Japanese
	word.English = RequestJson.Japanese
	database.CurrentUser(ctx, &user)

	word.UserID = user.ID

	database.CreateWord(word)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(400)
	}

}
package handler

import (
	"ej-ex3-backend/database"
	"ej-ex3-backend/model"
	"log"
	"strconv"

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
	word.English = RequestJson.English
	database.CurrentUser(ctx, &user)

	word.UserID = user.ID

	database.CreateWord(word, user)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(400)
		return
	}

}

func GetALLWordsByUser(ctx *gin.Context) {
	idstr := ctx.Param("id")

	id, err := strconv.Atoi(idstr)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(404)
		return
	}

	user, err := database.FindUserById(int64(id))

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(404)
		return
	}

	ctx.JSON(200, user.UserWords)
}

func GetALLWordsByUsers(ctx *gin.Context) {
	uw, err := database.GetALLUsersWords()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(200, uw)
}

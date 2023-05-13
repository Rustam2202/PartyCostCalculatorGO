package readers

import (
	"fmt"
	"net/http"

	"party-calc/internal"
	"party-calc/internal/person"
	"party-calc/utils"

	"github.com/gin-gonic/gin"
)

func JsonHandler(ctx *gin.Context) {
	var pers person.Persons
	err := ctx.ShouldBindJSON(&pers)
	if err != nil {
		utils.Logger.Error("Incorrect input JSON format")
		panic(err)
	}
	result := internal.CalculateDebts(pers, 1)
	ctx.JSON(http.StatusOK, result)
}

func StartServer() {
	router := gin.Default()
	router.GET("/", JsonHandler)
	err := router.Run(fmt.Sprintf(":%d", utils.Cfg.Port))
	if err != nil {
		utils.Logger.Error("Server couldn`t start")
		panic(err)
	}
}

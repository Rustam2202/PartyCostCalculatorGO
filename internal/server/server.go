package server

import (
	"fmt"
	"net/http"

	"party-calc/internal/person"
	"party-calc/internal/service"
	"party-calc/internal/logger"
	"party-calc/internal/config"

	"github.com/gin-gonic/gin"
)

func JsonHandler(ctx *gin.Context) {
	var pers person.Persons
	err := ctx.ShouldBindJSON(&pers)
	if err != nil {
		logger.Logger.Error("Incorrect input JSON format")
		return
	}
	result := service.CalculateDebts(pers)
	//ctx.JSON(http.StatusOK, result)
	ctx.JSON(http.StatusOK, result)
}

func StartServer() {
	router := gin.Default()
	router.POST("/", JsonHandler)
	err := router.Run(fmt.Sprintf(":%d", config.Cfg.Server.Port))
	if err != nil {
		logger.Logger.Error("Server couldn`t start")
		return
	}
}

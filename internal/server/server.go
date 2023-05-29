package server

import (
	"fmt"
	"net/http"

	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"
	"party-calc/internal/person"
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StartServer() {
	router := gin.Default()
	router.POST("/", JsonHandler)
	router.POST("/addper", AddPersonToDbHandler)
	router.GET("/getper", GetPersonFromDbHandler)
	err := router.Run(fmt.Sprintf(":%d", config.Cfg.Server.Port))
	if err != nil {
		logger.Logger.Error("Server couldn`t start:", zap.Error(err))
		return
	}
}

func JsonHandler(ctx *gin.Context) {
	var pers person.Persons
	err := ctx.ShouldBindJSON(&pers)
	if err != nil {
		logger.Logger.Error("Incorrect input JSON format")
		return
	}
	result := service.CalculateDebts(pers)
	ctx.JSON(http.StatusOK, result)
}

func AddPersonToDbHandler(ctx *gin.Context) {
	var db database.DataBase
	var per = models.Person{}
	name := ctx.Query("name")
	per.Name = name
	_, err := db.AddPerson(per)
	if err != nil {
		logger.Logger.Fatal("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func GetPersonFromDbHandler(ctx *gin.Context) {
	var db database.DataBase
	name := ctx.Query("name")
	_, err := db.GetPerson(name)
	if err != nil {
		logger.Logger.Fatal("couldn't GET person: ", zap.Error(err))
		return
	}
}

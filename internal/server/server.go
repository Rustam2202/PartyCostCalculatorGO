package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"
	"party-calc/internal/person"
	"party-calc/internal/server/config"
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StartServer() {
	var cfg config.ServerConfig
	cfg.LoadConfig()

	router := gin.Default()
	router.POST("/", JsonHandler)

	router.POST("/addPerson", AddPersonHandler)
	router.GET("/getPerson", GetPersonHandler)
	router.PUT("/updatePerson", UpdatePersonHandler)
	router.DELETE("/deletPerson", DeletePersonHandler)

	router.POST("/addEvent", AddEventHandler)
	router.GET("/getEvent", GetEventHandler)
	router.PUT("/updateEvent", UpdateEventHandler)
	router.DELETE("/deleteEvent", DeleteEventHandler)

	router.POST("/addPersonToEvent", AddPersonToEventHandler)
	router.POST("/addPersonToEventWithSpent", AddPersonToEventWithSpentHandler)
	router.GET("/getPersonFromEvents", GetPersEventsHandler)
	router.PUT("/updatePersonInEvents", UpdatePersEventsHandler)
	router.DELETE("/deletePersonInEvents", DeletePersonFromEventsHandler)

	err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port))
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

func AddPersonHandler(ctx *gin.Context) {
	var db database.DataBase
	var per = models.Person{}
	name := ctx.Query("name")
	per.Name = name
	id, err := db.AddPerson(per)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
	ctx.String(http.StatusOK, "", id)
}

func GetPersonHandler(ctx *gin.Context) {
	var db database.DataBase
	name := ctx.Query("name")
	_, err := db.GetPerson(name)
	if err != nil {
		//logger.Logger.Error("couldn't GET person: ", zap.Error(err))
		return
	}
}

func UpdatePersonHandler(ctx *gin.Context) {
	var db database.DataBase
	var per = models.Person{}
	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	per.Name = ctx.Query("name")
	err := db.UpdatePerson(id, per)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func DeletePersonHandler(ctx *gin.Context) {
	var db database.DataBase
	id, _ := strconv.Atoi(ctx.Query("id"))
	err := db.DeletePerson(int64(id))
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func AddEventHandler(ctx *gin.Context) {
	var db database.DataBase
	var ev = models.Event{}
	ev.Name = ctx.Query("name")
	date, err := time.Parse("2006-01-02", ctx.Query("date"))
	if err != nil {

	}
	ev.Date = date
	_, err = db.AddEvent(ev)
	if err != nil {

	}
}

func GetEventHandler(ctx *gin.Context) {
	var db database.DataBase
	name := ctx.Query("name")
	_, err := db.GetEvent(name)
	if err != nil {
		//logger.Logger.Error("couldn't GET person: ", zap.Error(err))
		return
	}
}

func UpdateEventHandler(ctx *gin.Context) {
	var db database.DataBase
	var ev = models.Event{}
	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	date, err := time.Parse("2006-01-02", ctx.Query("date"))
	if err != nil {

	}
	ev.Date = date
	ev.Name = ctx.Query("name")
	err = db.UpdateEvent(id, ev)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func DeleteEventHandler(ctx *gin.Context) {
	var db database.DataBase
	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	err := db.DeleteEvent(id)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func AddPersonToEventHandler(ctx *gin.Context) {
	var db database.DataBase
	perid, _ := strconv.ParseInt(ctx.Query("perid"), 10, 64)
	evid, _ := strconv.ParseInt(ctx.Query("evid"), 10, 64)
	err := db.AddPersonToEvent(perid, evid)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func AddPersonToEventWithSpentHandler(ctx *gin.Context) {
	var db database.DataBase
	perid, _ := strconv.ParseInt(ctx.Query("perid"), 10, 64)
	evid, _ := strconv.ParseInt(ctx.Query("evid"), 10, 64)
	spent, _ := strconv.ParseFloat(ctx.Query("perid"), 64)
	factor, _ := strconv.Atoi(ctx.Query("evid"))

	err := db.AddPersonToEventWithSpent(perid, evid, spent, factor)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func GetPersEventsHandler(ctx *gin.Context) {
	var db database.DataBase
	name := ctx.Query("name")
	_, err := db.GetPersEvents(name)
	if err != nil {
		//logger.Logger.Error("couldn't GET person: ", zap.Error(err))
		return
	}
}

func UpdatePersEventsHandler(ctx *gin.Context) {
	var db database.DataBase
	perid, _ := strconv.ParseInt(ctx.Query("perid"), 10, 64)
	evid, _ := strconv.ParseInt(ctx.Query("evid"), 10, 64)
	spent, _ := strconv.ParseFloat(ctx.Query("perid"), 64)
	factor, _ := strconv.Atoi(ctx.Query("evid"))

	err := db.UpdatePersEvents(perid, evid, spent, factor)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

func DeletePersonFromEventsHandler(ctx *gin.Context) {
	var db database.DataBase
	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)
	err := db.DeletePersonFromEvents(id)
	if err != nil {
		//logger.Logger.Error("couldn't INSERT person: ", zap.Error(err))
		return
	}
}

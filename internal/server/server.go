package server

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"
	"party-calc/internal/person"
	srvCfg "party-calc/internal/server/config"

	//dbCfg "party-calc/internal/database/config"
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var db database.DataBase

func StartServer() {
	var serverConfig srvCfg.ServerConfig
	//var databaseConfig dbCfg.DatabaseConfig

	srvCfgPath := flag.String("srvcfg", "./internal/server/config/", "path to server config file")
	dbCfgPath := flag.String("dbcfg", "./internal/database/config/", "path to database config file")
	flag.Parse()
	
	serverConfig.LoadConfig(*srvCfgPath)
	db.CFG.LoadConfig(*dbCfgPath)

	err := db.Open()
	if err != nil {
		logger.Logger.Error("Database couldn`t open:", zap.Error(err))
		return
	}
	defer db.DB.Close()

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

	err = router.Run(fmt.Sprintf(":%d", serverConfig.Server.Port))
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
	var per = models.Person{}
	name := ctx.Query("name")
	per.Name = name
	id, err := db.AddPerson(per)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added with id:": id})
}

func GetPersonHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	per, err := db.GetPerson(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

func UpdatePersonHandler(ctx *gin.Context) {
	var per = models.Person{}
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing id:": err})
		return
	}
	// id := ctx.GetInt64("id") // ?? returns 0
	per.Name = ctx.Query("name")
	err = db.UpdatePerson(id, per)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update person in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person updated:": id})
}

func DeletePersonHandler(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing id:": err})
		return
	}
	//id := ctx.GetInt64("id") // ?? returns 0
	err = db.DeletePerson(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted:": id})
}

func AddEventHandler(ctx *gin.Context) {
	var ev = models.Event{}
	ev.Name = ctx.Query("name")
	//ev.Date = ctx.GetTime("date") // ?? don't parsing
	date, err := time.Parse("2006-01-02", ctx.Query("date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	ev.Date = date
	id, err := db.AddEvent(ev)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with added event to database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event added with id:": id})
}

func GetEventHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	ev, err := db.GetEvent(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, ev)
}

func UpdateEventHandler(ctx *gin.Context) {
	var ev = models.Event{}
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing id:": err})
		return
	}
	date, err := time.Parse("2006-01-02", ctx.Query("date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	ev.Date = date
	ev.Name = ctx.Query("name")
	err = db.UpdateEvent(id, ev)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update event in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event updated:": id})
}

func DeleteEventHandler(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing id:": err})
		return
	}
	err = db.DeleteEvent(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event deleted:": id})
}

func AddPersonToEventHandler(ctx *gin.Context) {
	perid, err := strconv.ParseInt(ctx.Query("perid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing person id: ": err})
		return
	}
	evid, err := strconv.ParseInt(ctx.Query("evid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing event id: ": err})
		return
	}
	id, err := db.AddPersonToEvent(perid, evid)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to events: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added to events with id: ": id})
}

func AddPersonToEventWithSpentHandler(ctx *gin.Context) {
	perid, err := strconv.ParseInt(ctx.Query("personId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing person id: ": err})
		return
	}
	evid, err := strconv.ParseInt(ctx.Query("eventId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing event id: ": err})
		return
	}
	spent, err := strconv.ParseFloat(ctx.Query("spent"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing spent: ": err})
		return
	}
	factor, err := strconv.Atoi(ctx.Query("factor"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing factor: ": err})
		return
	}

	id, err := db.AddPersonToEventWithSpent(perid, evid, spent, factor)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added to events with id: ": id})
}

func GetPersEventsHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	perEv, err := db.GetPersEvents(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with getting person from events: ": err})
		return
	}
	ctx.JSON(http.StatusOK, perEv)
}

func UpdatePersEventsHandler(ctx *gin.Context) {
	perid, err := strconv.ParseInt(ctx.Query("personId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing person id: ": err})
		return
	}
	evid, err := strconv.ParseInt(ctx.Query("eventId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing event id: ": err})
		return
	}
	spent, err := strconv.ParseFloat(ctx.Query("spent"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing spent: ": err})
		return
	}
	factor, err := strconv.Atoi(ctx.Query("factor"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing factor: ": err})
		return
	}

	err = db.UpdatePersEvents(perid, evid, spent, factor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update person in events: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person in events updated: ": perid})
}

func DeletePersonFromEventsHandler(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing id:": err})
		return
	}
	err = db.DeletePersonFromEvents(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person in events: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted from events: ": id})
}

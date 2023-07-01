package handlers

import (
	"net/http"
	"party-calc/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(s *service.EventService) *EventHandler {
	return &EventHandler{service: s}
}

func (h *EventHandler) Add(ctx *gin.Context) {
	name := ctx.Query("name")
	//ev.Date = ctx.GetTime("date") // ?? don't parsing
	date, err := time.Parse("2006-01-02", ctx.Query("date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	 err = h.service.NewEvent(name, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with added event to database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event added with id:": ""})
}

func (h *EventHandler) Get(ctx *gin.Context) {
	name := ctx.Query("name")
	_, err := h.service.GetEventById(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, "")
}

func (h *EventHandler) Update(ctx *gin.Context) {
	name := ctx.Query("name")
	newName := ctx.Query("newname")
	newDate, err := time.Parse("2006-01-02", ctx.Query("newdate"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	err = h.service.UpdateEvent(name, newName, newDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update event in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event updated:": ""})
}

func (h *EventHandler) Delete(ctx *gin.Context) {
	//id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)

	name := ctx.Query("name")
	err := h.service.DeleteEvent(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event deleted:": ""})
}

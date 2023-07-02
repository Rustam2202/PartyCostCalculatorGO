package handlers

import (
	"net/http"
	"party-calc/internal/domain"
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
	req := struct {
		Name string `json:"name"`
		Date string `json:"date"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	id, err := h.service.NewEvent(req.Name, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with added event to database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event added with id:": id})
}

func (h *EventHandler) Get(ctx *gin.Context) {
	req := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	var ev *domain.Event
	if req.Id != 0 {
		ev, err = h.service.GetEventById(req.Id)
	} else if req.Name != "" {
		ev, err = h.service.GetEventByName(req.Name)
	} else {
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, *ev)
}

func (h *EventHandler) Update(ctx *gin.Context) {
	req := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Date string `json:"date"`
	}{}
	err := ctx.ShouldBindJSON(&req)

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	err = h.service.UpdateEvent(req.Id, req.Name, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update event in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event updated: ": req.Name})
}

func (h *EventHandler) Delete(ctx *gin.Context) {
	req := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}
	err := ctx.ShouldBindJSON(&req)

	if req.Id != 0 {
		err = h.service.DeleteEventById(req.Id)
	} else if req.Name != "" {
		err = h.service.DeleteEventByName(req.Name)
	} else {
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event deleted:": req.Id})
}

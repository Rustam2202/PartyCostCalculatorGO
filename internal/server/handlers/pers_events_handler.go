package handlers

import (
	"net/http"
	"party-calc/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersEventsHandler struct {
	service *service.PersonsEventsService
}

func NewPersEventsHandler(s *service.PersonsEventsService) *PersEventsHandler {
	return &PersEventsHandler{service: s}
}

func (h *PersEventsHandler) Add(ctx *gin.Context) {
	per := ctx.Query("person")
	ev := ctx.Query("event")
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
	id, err := h.service.AddPersonToPersEvent(per, ev, spent, factor)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added to events with id: ": id})
}

func (h *PersEventsHandler) Get(ctx *gin.Context) {
	name := ctx.Query("name")
	_, err := h.service.GetPerson(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with getting person from events: ": err})
		return
	}
	ctx.JSON(http.StatusOK, "")
}

func (h *PersEventsHandler) Update(ctx *gin.Context) {
	per := ctx.Query("person")
	ev := ctx.Query("event")
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
	err = h.service.UpdatePerson(per, ev, spent, factor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update person in events: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person in events updated: ": ""})
}

func (h *PersEventsHandler) Delete(ctx *gin.Context) {
	name := ctx.Query("name")
	err := h.service.DeletePerson(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person in events: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted from events: ": ""})
}

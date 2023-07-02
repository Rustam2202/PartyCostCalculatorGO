package handlers

import (
	"net/http"
	"party-calc/internal/domain"
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
)

type PersEventsHandler struct {
	service *service.PersonsEventsService
}

func NewPersEventsHandler(s *service.PersonsEventsService) *PersEventsHandler {
	return &PersEventsHandler{service: s}
}

func (h *PersEventsHandler) Add(ctx *gin.Context) {
	req := struct {
		PerId  int64   `json:"person_id"`
		EvId   int64   `json:"event_id"`
		Spent  float64 `json:"spent"`
		Factor int     `json:"factor"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Incorrect request: ": err})
		return
	}
	if req.Factor == 0 {
		req.Factor = 1
	}
	id, err := h.service.AddPersonToPersEvent(req.PerId, req.EvId, req.Spent, req.Factor)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Failed Insert to 'persons_events' table: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added to events with id: ": id})
}

func (h *PersEventsHandler) Get(ctx *gin.Context) {
	req := struct {
		PerId int64 `json:"person_id"`
		EvId  int64 `json:"event_id"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	var ev *domain.PersonsAndEvents
	if req.PerId != 0 {
		ev, err = h.service.GetByPersonId(req.PerId)
	} else if req.EvId != 0 {
		ev, err = h.service.GetByEventId(req.EvId)
	} else {
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, *ev)
}

func (h *PersEventsHandler) Update(ctx *gin.Context) {
	req := struct {
		Id     int64   `json:"id"`
		PerId  int64   `json:"person_id"`
		EvId   int64   `json:"event_id"`
		Spent  float64 `json:"spent"`
		Factor int     `json:"factor"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	err = h.service.Update(req.Id, req.PerId, req.EvId, req.Spent, req.Factor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update event in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event updated:": ""})
}

func (h *PersEventsHandler) Delete(ctx *gin.Context) {
	req := struct {
		Id int64 `json:"id"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if req.Id != 0 {
		err = h.service.Delete(req.Id)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event deleted:": ""})
}

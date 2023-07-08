package events

import (
	"net/http"
	"party-calc/internal/domain"

	"github.com/gin-gonic/gin"
)

type GetEventRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (h *EventHandler) Get(ctx *gin.Context) {
	var req GetEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request:": err})
		return
	}
	var ev *domain.Event
	if req.Id != 0 {
		ev, err = h.service.GetEventById(ctx, req.Id)
	} else if req.Name != "" {
		ev, err = h.service.GetEventByName(ctx, req.Name)
	} else {
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, *ev)
}

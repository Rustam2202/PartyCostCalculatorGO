package personsevents

import (
	"net/http"
	"party-calc/internal/domain"

	"github.com/gin-gonic/gin"
)

type GetPersonEventRequest struct {
	PerId int64 `json:"person_id"`
	EvId  int64 `json:"event_id"`
}

func (h *PersEventsHandler) Get(ctx *gin.Context) {
	var req GetPersonEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request: ": err})
		return
	}
	var ev *domain.PersonsAndEvents
	if req.PerId != 0 {
		ev, err = h.service.GetByPersonId(ctx, req.PerId)
	} else if req.EvId != 0 {
		ev, err = h.service.GetByEventId(ctx, req.EvId)
	} else {
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, *ev)
}

package personsevents

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePersonEventRequest struct {
	Id     int64   `json:"id"`
	PerId  int64   `json:"person_id"`
	EvId   int64   `json:"event_id"`
	Spent  float64 `json:"spent"`
	Factor int     `json:"factor"`
}

func (h *PersEventsHandler) Update(ctx *gin.Context) {
	var req UpdatePersonEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request: ": err})
		return
	}
	err = h.service.Update(ctx, req.Id, req.PerId, req.EvId, req.Spent, req.Factor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update event in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event updated:": ""})
}

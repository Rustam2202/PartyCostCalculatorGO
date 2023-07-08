package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteEventRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (h *EventHandler) Delete(ctx *gin.Context) {
	var req DeleteEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request:": err})
		return
	}
	if req.Id != 0 {
		err = h.service.DeleteEventById(ctx, req.Id)
	} else if req.Name != "" {
		err = h.service.DeleteEventByName(ctx, req.Name)
	} else {
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event deleted:": req.Id})
}

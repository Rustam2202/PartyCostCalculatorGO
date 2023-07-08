package personsevents

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeletePersonEventRequest struct {
	Id int64 `json:"id"`
}

func (h *PersEventsHandler) Delete(ctx *gin.Context) {
	var req DeletePersonEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request: ": err})
		return
	}
	if req.Id != 0 {
		err = h.service.Delete(ctx, req.Id)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete event from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event deleted:": ""})
}

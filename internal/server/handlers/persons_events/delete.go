package personsevents

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeletePersonEventRequest struct {
	Id int64 `json:"id"`
}

//	@Summary		Delete a person-event
//	@Description	Delete a record of peson existed in event by Id from database
//	@Tags			Person-Event
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DeletePersonEventRequest	true	"Delete Person-Event Request"
//	@Success		200		{object}	int64
//	@Failure		304		{object}	handlers.ErrorResponce
//	@Failure		400		{object}	handlers.ErrorResponce
//	@Router			/persEvents [delete]
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

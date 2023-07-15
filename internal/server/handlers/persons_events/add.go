package personsevents

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddPersonEventRequest struct {
	PerId  int64   `json:"person_id"`
	EvId   int64   `json:"event_id"`
	Spent  float64 `json:"spent"`
	Factor int     `json:"factor"`
}

//	@Summary		Add a person-event
//	@Description	Add a new record of peson existed in event by Id to database
//	@Tags			Person-Event
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AddPersonEventRequest	true	"Add Person-Event Request"
//	@Success		200		{object}	int64
//	@Failure		304		{object}	handlers.ErrorResponce
//	@Failure		400		{object}	handlers.ErrorResponce
//	@Router			/persEvents [post]
func (h *PersEventsHandler) Add(ctx *gin.Context) {
	var req AddPersonEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request: ": err})
		return
	}
	if req.Factor == 0 {
		req.Factor = 1
	}
	id, err := h.service.AddPersonToPersEvent(ctx, req.PerId, req.EvId, req.Spent, req.Factor)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Failed Insert to 'persons_events' table: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added to events with id: ": id})
}

package personsevents

import (
	"net/http"
	"party-calc/internal/domain"
	"party-calc/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

type AddPersonEventRequest struct {
	PerId  int64   `json:"person_id" default:"123456789"`
	EvId   int64   `json:"event_id" default:"987654321"`
	Spent  float64 `json:"spent" default:"10.25"`
	Factor int     `json:"factor" default:"1"`
}

// @Summary		Add a person-event
// @Description	Add a new record of peson existed in event to database
// @Tags			Person-Event
// @Accept			json
// @Produce		json
// @Param			request	body		AddPersonEventRequest	true	"Add Person-Event Request"
// @Success		201		{object}	domain.PersonsAndEvents
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/persEvents [post]
func (h *PersEventsHandler) Add(ctx *gin.Context) {
	var req AddPersonEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	id, err := h.service.AddPersonToPersEvent(ctx, req.PerId, req.EvId, req.Spent, req.Factor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person-event data to database", Error: err})
		return
	}
	ctx.JSON(http.StatusCreated, domain.PersonsAndEvents{
		Id:       id,
		PersonId: req.PerId,
		EventId:  req.EvId,
		Spent:    req.Spent,
		Factor:   req.Factor})
}

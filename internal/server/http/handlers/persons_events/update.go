package personsevents

import (
	"net/http"
	"party-calc/internal/server/http/handlers"

	"github.com/gin-gonic/gin"
)

type UpdatePersonEventRequest struct {
	Id     int64   `json:"id"`
	PerId  int64   `json:"person_id"`
	EvId   int64   `json:"event_id"`
	Spent  float64 `json:"spent"`
	Factor int     `json:"factor" default:"1"`
}

// @Summary		Update a person-event
// @Description	Update a record of peson-event data
// @Tags			Person-Event
// @Accept			json
// @Produce		json
// @Param			request	body		UpdatePersonEventRequest	true	"Update Person-Event Request"
// @Success		200		{object}	UpdatePersonEventRequest
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/persEvents [put]
func (h *PersEventsHandler) Update(ctx *gin.Context) {
	var req UpdatePersonEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.Update(ctx, req.Id, req.PerId, req.EvId, req.Spent, req.Factor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to update a person-event data in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, req)
}

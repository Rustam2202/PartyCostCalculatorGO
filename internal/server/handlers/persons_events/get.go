package personsevents

import (
	"net/http"
	"party-calc/internal/server/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary		Get a persons-event
// @Description	Get an array of peson-event records by EventId
// @Tags			Person-Event
// @Accept			json
// @Produce		json
// @Param			event_id     path    int     true        "Event Id"
// @Success		200		{object}	[]domain.PersonsAndEvents
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/persEvents/{event_id} [get]
func (h *PersEventsHandler) Get(ctx *gin.Context) {
	req := ctx.Param("event_id")
	perId, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	ev, err := h.service.GetByEventId(ctx, perId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a persons-event data from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, ev)
}

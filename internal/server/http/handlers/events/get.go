package events

import (
	"net/http"
	"party-calc/internal/server/http/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary		Get a event
// @Description	Get a event from database
// @Tags			Event
// @Accept			json
// @Produce		json
// @Param			id     path    int     true        "Event Id"
// @Success		200		{object}	domain.Event
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/event/{id} [get]
func (h *EventHandler) Get(ctx *gin.Context) {
	req := ctx.Param("id")
	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	ev, err := h.service.GetEventById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, ev)
}

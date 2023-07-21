package events

import (
	"net/http"
	"party-calc/internal/server/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary		Delete a event
// @Description	Delete a event from database
// @Tags			Event
// @Accept			json
// @Produce		json
// @Param			id     path    int     true        "Event Id"
// @Success		200
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/event/{id} [delete]
func (h *EventHandler) Delete(ctx *gin.Context) {
	req := ctx.Param("id")
	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.DeleteEventById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to delete a event from database", Error: err})
		return
	}
	ctx.Status(http.StatusOK)
}

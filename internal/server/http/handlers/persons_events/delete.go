package personsevents

import (
	"net/http"
	"party-calc/internal/server/http/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Delete a person-event
//	@Description	Delete a record of peson existed in event by Id from database
//	@Tags			Person-Event
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Event Id"
//	@Success		200
//	@Failure		400	{object}	handlers.ErrorResponce
//	@Failure		500	{object}	handlers.ErrorResponce
//	@Router			/persEvents/{id} [delete]
func (h *PersEventsHandler) Delete(ctx *gin.Context) {
	req := ctx.Param("id")
	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to delete a person-event data from database", Error: err})
		return
	}
	ctx.Status(http.StatusOK)
}

package persons

import (
	"net/http"
	"party-calc/internal/server/http/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Delete a person
//	@Description	Delete a person from database
//	@Tags			Person
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Person Id"
//	@Success		200
//	@Failure		400	{object}	handlers.ErrorResponce
//	@Failure		500	{object}	handlers.ErrorResponce
//	@Router			/person/{id} [delete]
func (h *PersonHandler) Delete(ctx *gin.Context) {
	req := ctx.Param("id")
	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.DeletePersonById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to delete a person from database", Error: err})
		return
	}
	ctx.Status(http.StatusOK)
}

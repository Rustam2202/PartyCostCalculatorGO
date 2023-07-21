package persons

import (
	"net/http"
	"party-calc/internal/server/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary		Get a person
// @Description	Get a person from database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param		id     path    int     true        "Person Id"
// @Success		200		{object}	domain.Person
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person/{id} [get]
func (h *PersonHandler) Get(ctx *gin.Context) {
	req := ctx.Param("id")
	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	per, err := h.service.GetPersonById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

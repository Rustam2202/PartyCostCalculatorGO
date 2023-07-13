package persons

import (
	"net/http"
	"party-calc/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

type AddPersonRequest struct {
	Name string `json:"name"`
}

type AddPersonResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary Add a person
// @Description Add a new person to database
// @Tags Person
// @Accept json
// @Produce json
// @Param request body AddPersonRequest true "Add Person Request"
// @Success 200 {object} AddPersonResponse
// @Failure 304 {object} handlers.ErrorResponce
// @Failure 400 {object} handlers.ErrorResponce
// @Router /person [post]
func (h *PersonHandler) Add(ctx *gin.Context) {
	var req AddPersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed parse request", Error: err})
		return
	}
	id, err := h.service.NewPerson(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusNotModified,
			handlers.ErrorResponce{Message: "Failed add new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, AddPersonResponse{Id: id, Name: req.Name})
}

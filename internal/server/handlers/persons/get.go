package persons

import (
	"net/http"
	"party-calc/internal/domain"
	"party-calc/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

type GetPersonRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary Get a person
// @Description Get a person from database
// @Tags Person
// @Accept json
// @Produce json
// @Param request body GetPersonRequest true "Get Person Request"
// @Success 200 {object} domain.Person
// @Failure 400 {object} handlers.ErrorResponce
// @Failure 500 {object} handlers.ErrorResponce
// @Router /person [get]
func (h *PersonHandler) GetById(ctx *gin.Context) {
	var req GetPersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, 
			handlers.ErrorResponce{Message: "Failed parse request", Error: err})
		return
	}
	var per *domain.Person
	if req.Id != 0 {
		per, err = h.service.GetPersonById(ctx, req.Id)
	} else if req.Name != "" {
		per, err = h.service.GetPersonByName(ctx, req.Name)
	} else {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Id=0 or empty Name in request", Error: err})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Error with getting person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

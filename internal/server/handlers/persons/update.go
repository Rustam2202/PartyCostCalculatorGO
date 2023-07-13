package persons

import (
	"net/http"
	"party-calc/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

type UpdatePersonRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary Update a person
// @Description Update a person in database
// @Tags Person
// @Accept json
// @Produce json
// @Param request body UpdatePersonRequest true "Update Person Request"
// @Success 200 {object} domain.Person
// @Failure 400 {object} handlers.ErrorResponce
// @Failure 500 {object} handlers.ErrorResponce
// @Router /person [put]
func (h *PersonHandler) Update(ctx *gin.Context) {
	var req UpdatePersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed parse request", Error: err})
		return
	}
	err = h.service.UpdatePerson(ctx, req.Id, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Error with update person in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person updated: ": req.Name})
}

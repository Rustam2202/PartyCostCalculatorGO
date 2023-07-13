package persons

import (
	"net/http"
	"party-calc/internal/domain"

	"github.com/gin-gonic/gin"
)

type GetPersonRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary Get a Person
// @Description get a Person model by ID or NAME
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Person
// @Router /person [get]
func (h *PersonHandler) GetById(ctx *gin.Context) {
	var req GetPersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request:": err})
		return
	}
	var per *domain.Person
	if req.Id != 0 {
		per, err = h.service.GetPersonById(ctx, req.Id)
	} else if req.Name != "" {
		per, err = h.service.GetPersonByName(ctx, req.Name)
	} else {
		ctx.JSON(http.StatusBadRequest, "Id=0 or empty Name in request")
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

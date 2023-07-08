package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePersonRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (h *PersonHandler) Update(ctx *gin.Context) {
	var req UpdatePersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request:": err})
		return
	}
	err = h.service.UpdatePerson(ctx, req.Id, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update person in database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person updated: ": req.Name})
}

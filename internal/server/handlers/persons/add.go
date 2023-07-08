package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddPersonRequest struct {
	Name string `json:"name"`
}

func (h *PersonHandler) Add(ctx *gin.Context) {
	var req AddPersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request:": err})
		return
	}
	id, err := h.service.NewPerson(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added with id: ": id})
}

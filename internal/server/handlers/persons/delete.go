package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeletePersonRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary Delete a Person
// @Description get string by ID
// @Accept  json
// @Produce  json
// Success 200 {object} int64
// @Router /person [delete]
func (h *PersonHandler) Delete(ctx *gin.Context) {
	var req DeletePersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request:": err})
		return
	}
	if req.Id != 0 {
		err = h.service.DeletePersonById(ctx, req.Id)
	} else if req.Name != "" {
		err = h.service.DeletePersonByName(ctx, req.Name)
	} else {
		ctx.JSON(http.StatusBadRequest, "Id=0 or empty Name in request")
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted: ": req.Id})
}

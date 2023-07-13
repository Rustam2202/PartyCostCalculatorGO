package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeletePersonRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// @Summary Delete a person
// @Description Delete a person from database
// @Tags Person
// @Accept json
// @Produce json
// @Param request body DeletePersonRequest true "Delete Person Request"
// @Success 200 {object} int64
// @Failure 304 {object} handlers.ErrorResponce
// @Failure 400 {object} handlers.ErrorResponce
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

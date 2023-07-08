package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeletePersonRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

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
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Parsing JSON request error: ": err})
			return
		}
	}
	err = h.service.DeletePersonById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted: ": req.Id})
}

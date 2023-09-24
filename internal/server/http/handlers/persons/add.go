package persons

import (
	"net/http"
	"party-calc/internal/domain"
	"party-calc/internal/server/http/handlers"

	"github.com/gin-gonic/gin"
)

type AddPersonRequest struct {
	Name string `json:"name" default:"Some Person name"`
}

//	@Summary		Add a person
//	@Description	Add a new person to database
//	@Tags			Person
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AddPersonRequest	true	"Add Person Request"
//	@Success		201		{object}	domain.Person
//	@Failure		400		{object}	handlers.ErrorResponce
//	@Failure		500		{object}	handlers.ErrorResponce
//	@Router			/person [post]
func (h *PersonHandler) Add(ctx *gin.Context) {
	var req AddPersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	id, err := h.service.NewPerson(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusCreated, domain.Person{Id: id, Name: req.Name})
}

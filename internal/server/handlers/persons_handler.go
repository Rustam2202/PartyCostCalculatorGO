package handlers

import (
	"net/http"
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	service *service.PersonService
}

func NewPersonHandler(s *service.PersonService) *PersonHandler {
	return &PersonHandler{service: s}
}

func (h *PersonHandler) Add(ctx *gin.Context) {
	name := ctx.Query("name")
	id, err := h.service.NewPerson(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added with id: ": id})
}

func (h *PersonHandler) Get(ctx *gin.Context) {
	name := ctx.Query("name")
	per, err := h.service.GetPerson(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

func (h *PersonHandler) Update(ctx *gin.Context) {
	name := ctx.Query("name")
	newName := ctx.Query("newname")
	err := h.service.UpdatePerson(name, newName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update person in database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person updated: ": name})
}

func (h *PersonHandler) Delete(ctx *gin.Context) {
	name := ctx.Query("name")
	err := h.service.DeletePerson(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted: ": name})
}

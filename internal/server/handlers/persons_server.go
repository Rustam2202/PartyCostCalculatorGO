package server

import (
	"net/http"
	"party-calc/internal/database/models"
	"party-calc/internal/service"
	"strconv"

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
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added with id:": id})
}

func (h *PersonHandler)GetPersonHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	per, err := h.service.(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

func UpdatePersonHandler(ctx *gin.Context) {
	var per = models.Person{}
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing id:": err})
		return
	}
	// id := ctx.GetInt64("id") // ?? returns 0
	per.Name = ctx.Query("name")
	err = db.UpdatePerson(id, per)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update person in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person updated:": id})
}

func DeletePersonHandler(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with parsing id:": err})
		return
	}
	//id := ctx.GetInt64("id") // ?? returns 0
	err = db.DeletePerson(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person from database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted:": id})
}

package handlers

import (
	"net/http"
	"party-calc/internal/domain"
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
	req := struct {
		Name string `json:"name"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	id, err := h.service.NewPerson(req.Name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{"Error with added person to database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person added with id: ": id})
}

func (h *PersonHandler) Get(ctx *gin.Context) {
	req := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	var per *domain.Person
	if req.Id != 0 {
		per, err = h.service.GetPersonById(req.Id)
	} else if req.Name != "" {
		per, err = h.service.GetPersonByName(req.Name)
	} else {
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

func (h *PersonHandler) Update(ctx *gin.Context) {
	per := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}
	err := ctx.ShouldBindJSON(&per)

	err = h.service.UpdatePerson(per.Id, per.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update person in database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person updated: ": per.Name})
}

func (h *PersonHandler) Delete(ctx *gin.Context) {
	per := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}
	err := ctx.ShouldBindJSON(&per)
	if per.Id != 0 {
		err = h.service.DeletePersonById(per.Id)
	} else if per.Name != "" {
		err = h.service.DeletePersonByName(per.Name)
	} else {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Parsing JSON request error: ": err})
			return
		}
	}
	err = h.service.DeletePersonById(per.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with delete person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Person deleted: ": per.Id})
}

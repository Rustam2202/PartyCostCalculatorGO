package handlers

import (
	"net/http"
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
)

type CalcHandler struct {
	service *service.CalcService
}

func NewCalcHandler(s *service.CalcService) *CalcHandler {
	return &CalcHandler{service: s}
}

func (s *CalcHandler) GetEvent(ctx *gin.Context) {
	req := struct {
		EventId int64 `json:"event_id"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	result, err := s.service.CalcEvent(ctx, req.EventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

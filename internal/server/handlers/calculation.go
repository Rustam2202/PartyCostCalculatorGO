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

type GetEventRequest struct {
	EventId int64 `json:"event_id"`
}

func (s *CalcHandler) GetEvent(ctx *gin.Context) {
	var req GetEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Failed with request: ": err})
		return
	}
	result, err := s.service.CalculateEvent(ctx, req.EventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

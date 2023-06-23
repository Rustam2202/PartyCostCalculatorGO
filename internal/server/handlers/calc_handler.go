package handlers

import (
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
)

type CalcHandler struct {
	service *service.CalcService
}

func NewCalcHandler(s *service.CalcService) *CalcHandler {
	return &CalcHandler{service: s}
}

func (s *CalcHandler) GetPerson(ctx *gin.Context) {

}

func (s *CalcHandler) GetEvent(ctx *gin.Context) {

}

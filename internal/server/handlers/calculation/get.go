package calculation

import (
	"net/http"
	"party-calc/internal/server/handlers"
	"party-calc/internal/service"

	"github.com/gin-gonic/gin"
)

type CalcHandler struct {
	service *service.CalcService
}

func NewCalcHandler(s *service.CalcService) *CalcHandler {
	return &CalcHandler{service: s}
}

type CalculateRequest struct {
	EventId   int64   `json:"event_id" default:"987654321"`
	RoundRate float64 `json:"round_rate" default:"1.0"`
}

// @Summary		Calculate event data by Id
// @Description
// @Tags			Calculate
// @Accept			json
// @Produce		json
// @Param			request	body		CalculateRequest	true	"Calculate Event Request"
// @Success		200		{object}	service.EventData
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/calcEvent [get]
func (s *CalcHandler) GetEvent(ctx *gin.Context) {
	var req CalculateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed parse request", Error: err})
		return
	}
	result, err := s.service.CalculateEvent(ctx, req.EventId, req.RoundRate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to calculate event data", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

package calculate

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
	EventId int64 `json:"event_id"`
}

// @Summary		Calculate event data by Id
// @Description	Calculate
// @Tags			Calculate
// @Accept			json
// @Produce		json
// @Param			request	body		CalculateRequest	true	"Calculate Event Request"
// @Success		200		{object}	service.EventData
// @Failure		304		{object}	handlers.ErrorResponce
// @Failure		400		{object}	handlers.ErrorResponce
// @Router			/calcEvent [get]
func (s *CalcHandler) GetEvent(ctx *gin.Context) {
	var req CalculateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed parse request", Error: err})
		return
	}
	result, err := s.service.CalculateEvent(ctx, req.EventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with getting person from database: ": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

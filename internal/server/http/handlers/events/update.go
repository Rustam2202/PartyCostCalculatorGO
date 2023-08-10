package events

import (
	"net/http"
	"party-calc/internal/server/http/handlers"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateEventRequest struct {
	Id   int64  `json:"id" default:"9876543212"`
	Name string `json:"name" default:"Some new Event name"`
	Date string `json:"date" default:"2020-11-30"`
}

// @Summary		Update a event
// @Description	Update a event in database
// @Tags			Event
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateEventRequest	true	"Update Event Request"
// @Success		200		{object}	UpdateEventRequest
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/event [put]
func (h *EventHandler) Update(ctx *gin.Context) {
	var req UpdateEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse date", Error: err})
		return
	}
	err = h.service.UpdateEvent(ctx, req.Id, req.Name, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to update a event in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, req)
}

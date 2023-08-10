package events

import (
	"net/http"
	"party-calc/internal/domain"
	"party-calc/internal/server/http/handlers"
	"time"

	"github.com/gin-gonic/gin"
)

type AddEventRequest struct {
	Name string `json:"name" default:"Some Event name"`
	Date string `json:"date" default:"2020-12-31"`
}

// @Summary		Add a event
// @Description	Add a new event to database
// @Tags			Event
// @Accept			json
// @Produce		json
// @Param			request	body		AddEventRequest	true	"Add Event Request"
// @Success		201		{object}	domain.Event
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/event [post]
func (h *EventHandler) Add(ctx *gin.Context) {
	var req AddEventRequest
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
	id, err := h.service.NewEvent(ctx, req.Name, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusCreated, domain.Event{Id: id, Name: req.Name, Date: date})
}

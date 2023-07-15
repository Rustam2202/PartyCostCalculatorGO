package events

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddEventRequest struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type AddEventResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Date time.Time `json:"date"`
}

//	@Summary		Add a event
//	@Description	Add a new event to database
//	@Tags			Event
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AddEventRequest	true	"Add Event Request"
//	@Success		200		{object}	AddEventResponse
//	@Failure		304		{object}	handlers.ErrorResponce
//	@Failure		400		{object}	handlers.ErrorResponce
//	@Router			/event [post]
func (h *EventHandler) Add(ctx *gin.Context) {
	var req AddEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Failed parse request:": err})
		return
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with date parsing :": err})
		return
	}
	id, err := h.service.NewEvent(ctx, req.Name, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with added event to database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event added with id:": id})
}

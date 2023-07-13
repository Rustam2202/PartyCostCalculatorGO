package events

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateEventRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

// @Router /person [post]
func (h *EventHandler) Update(ctx *gin.Context) {
	var req UpdateEventRequest
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
	err = h.service.UpdateEvent(ctx, req.Id, req.Name, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error with update event in database:": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Event updated: ": req.Name})
}

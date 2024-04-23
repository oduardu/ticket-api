package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oduardu/ticket-api/schemas"
)

func ListTicketsHandler(ctx *gin.Context) {
	ticekts := []schemas.Ticket{}

	if err := db.Find(&ticekts).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-tickets", ticekts)
}

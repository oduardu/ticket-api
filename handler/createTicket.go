package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oduardu/ticket-api/schemas"
)

func CreateTicketHandler(ctx *gin.Context) {
	req := CreateTicketRequest{}

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		fmt.Errorf("validator error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ticket := schemas.Ticket{
		FullName: req.FullName,
		Phone:    req.Phone,
		Mail:     req.Mail,
		Category: req.Category,
		Message:  req.Message,
	}

	if err := db.Create(&ticket).Error; err != nil {
		fmt.Errorf("error create ticket: %v", err.Error())
		return
	}

	sendSuccess(ctx, "create-ticket", ticket)
}

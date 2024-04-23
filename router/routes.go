package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oduardu/ticket-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/create", handler.CreateTicketHandler)
		v1.GET("/list", handler.ListTicketsHandler)
	}
}

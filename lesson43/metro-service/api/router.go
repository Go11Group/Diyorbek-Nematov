package api

import (
	"metro-service/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(handle *handler.Handler) *gin.Engine {
	r := gin.Default() 

	card := r.Group("/cards")
	{
		card.GET("/", handle.GetCardsHandler)
		card.POST("/", handle.CreateCardHandler)
		card.GET("/:id", handle.GetCardByIDHandler)
		card.PUT("/:id", handle.UpdateCardHandler)
		card.DELETE("/:id", handle.DeleteCardHandler)
	}

	terminal := r.Group("/terminals")
	{
		terminal.GET("/", handle.GetTerminalsHandler)
		terminal.POST("/", handle.CreateTerminalHandler)
		terminal.GET("/:id", handle.GetTerminalByIDHandler)
		terminal.PUT("/:id", handle.UpdateTerminalHandler)
		terminal.DELETE("/:id", handle.DeleteTerminalHandler)
	}

	station := r.Group("/stations")
	{
		station.GET("/", handle.GetStationsHandler)
		station.POST("/", handle.CreateStationHandler)
		station.GET("/:id", handle.GetStationByIDHandler)
		station.PUT("/:id", handle.UpdateStationHandler)
		station.DELETE("/:id", handle.DeleteStationHandler)
	}

	transaction := r.Group("/transactions")
	{
		transaction.GET("/", handle.GetTransactionsHandler)
		transaction.POST("/", handle.CreateTransactionHandler)
		transaction.GET("/:id", handle.GetTransactionByIDHandler)
		transaction.PUT("/:id", handle.UpdateTransactionHandler)
		transaction.DELETE("/:id", handle.DeleteTransactionHandler)
	}

	r.GET("/check_balance", handle.CheckBalance)
	r.GET("/balance", handle.GetBalance)	

	return r
}
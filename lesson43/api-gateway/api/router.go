package api

import (
	"api-gateway/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(handle handler.Handler) *gin.Engine {
	r := gin.Default()

	user := r.Group("/users")
	{
		user.POST("/", handle.CreateUser)
		user.GET("/", handle.GetUsers)
		user.GET("/:id", handle.GetUser)
		user.PUT("/:id", handle.UpdateUser)
		user.DELETE("/:id", handle.DeleteUser)
	}

	cards := r.Group("/cards")
	{
		cards.POST("/", handle.CreateCard)
		cards.GET("/:id", handle.GetCard)
		cards.GET("/", handle.GetCards)
		cards.PUT("/:id", handle.UpdateCard)
		cards.DELETE("/:id", handle.DeleteCard)
	}

	station := r.Group("/stations")
	{
		station.POST("/", handle.CreateStation)
		station.GET("/:id", handle.GetStation)
		station.GET("/", handle.GetStations)
		station.PUT("/:id", handle.UpdateStation)
		station.DELETE("/:id", handle.DeleteStation)
	}

	terminal := r.Group("/terminals")
	{
		terminal.POST("/", handle.CreateTerminal)
		terminal.GET("/:id", handle.GetTerminal)
		terminal.GET("/", handle.GetTerminals)
		terminal.PUT("/:id", handle.UpdateTerminal)
		terminal.DELETE("/:id", handle.DeleteTerminal)
	}

	transaction := r.Group("/transactions")
	{
		transaction.POST("/", handle.CreateTransaction)
		transaction.GET("/:id", handle.GetTransaction)
		transaction.GET("/", handle.GetTransactions)
		transaction.PUT("/:id", handle.UpdateTransaction)
		transaction.DELETE("/:id", handle.DeleteTransaction)
	}

	r.GET("/check_balance", handle.CheckBalance)
	r.GET("/balance", handle.GetBalance)

	return r
}
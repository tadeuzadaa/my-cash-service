package transport

import (
	"github.com/gin-gonic/gin"
	"my-cash-service/internal/transport/handlers"
)

func SetupRoutes(r *gin.RouterGroup) {
	r.POST("/mcs/createUser", handlers.CreateUser)
	r.POST("/mcs/login", handlers.LoginUser)
	r.GET("/mcs/user/:id", handlers.GetUserById)
	r.PUT("/mcs/user/:id", handlers.UpdateUser)
	r.DELETE("/mcs/user/:id", handlers.DeleteUser)

	r.POST("/mcs/transactions", handlers.CreateTransaction)
	r.GET("/mcs/transactions/:userId", handlers.GetTransactionsByUserId)
	r.PUT("/mcs/transactions/:id", handlers.UpdateTransaction)
	r.DELETE("/mcs/transactions/:id", handlers.DeleteTransaction)
}

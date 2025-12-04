package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"my-cash-service/internal/infra/database"
	"my-cash-service/internal/transport"
)

func main() {
	// Conecta ao banco de dados
	conn := database.ConnectDB()
	defer conn.Close(context.Background())

	// Configura o servidor de rotas Gin
	r := gin.Default()
	transport.SetupRoutes(&r.RouterGroup)

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}

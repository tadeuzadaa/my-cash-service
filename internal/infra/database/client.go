package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	url := "postgres://postgres:11650058@localhost:1165/my-cash-service-db"

	// cria a conexão
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	fmt.Println("Conexão bem-sucedida com o Postgres!")
	return conn
}

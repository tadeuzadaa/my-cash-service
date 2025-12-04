package core

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Transaction struct {
	ID       int             `json:"id"`
	UserID   int             `json:"user_id"`
	Type     TransactionType `json:"type"`
	Category string          `json:"category"`
	Value    float64         `json:"value"`
	Date     time.Time       `json:"date"`
	Desc     string          `json:"description"`
}

type TransactionType string

const (
	Receita TransactionType = "RECEITA"
	Despesa TransactionType = "DESPESA"
)

package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-cash-service/internal/core"
	"net/http"
)

var (
	userid = "userId"
	trID   = "id"
)

func CreateTransaction(c *gin.Context) {
	var t core.Transaction

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if t.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is missing!"})
		return
	}

	_, err := DBconn.Exec(
		c,
		"INSERT INTO transactions (user_id, type, category, value, date, description) VALUES ($1, $2, $3, $4, $5, $6)",
		t.UserID,
		t.Type,
		t.Category,
		t.Value,
		t.Date,
		t.Desc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar transação"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transação criada com sucesso!"})
}

func GetTransactionsByUserId(c *gin.Context) {
	userId := c.Param(userid)

	var transactions []core.Transaction
	var t core.Transaction

	rows, err := DBconn.Query(
		c,
		"SELECT id, user_id, type, category, value, date, description FROM transactions WHERE user_id=$1",
		userId,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar transações"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.UserID, &t.Type, &t.Category, &t.Value, &t.Date, &t.Desc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler transação"})
			return
		}
		transactions = append(transactions, t)
	}

	c.JSON(http.StatusOK, transactions)
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param(trID)

	var t core.Transaction

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tags, err := DBconn.Exec(
		c,
		"UPDATE transactions SET type=$1, category=$2, value=$3, date=$4, description=$5 WHERE id=$6",
		t.Type,
		t.Category,
		t.Value,
		t.Date,
		t.Desc,
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar transação"})
		return
	}
	if tags.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transação não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transação atualizada com sucesso!"})

}

func DeleteTransaction(c *gin.Context) {
	userId := c.Param(userid)
	id := c.Param(trID)

	var t core.Transaction

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tags, err := DBconn.Exec(
		c,
		"DELETE FROM transactions WHERE id=$1 and user_id=$2",
		id,
		userId,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar transação"})
		return
	}

	if tags.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

}

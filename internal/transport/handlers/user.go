package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"my-cash-service/internal/core"
	"my-cash-service/internal/infra/database"
	"net/http"
)

var (
	user   core.User
	DBconn = database.ConnectDB()
	id     = "id"
)

func CreateUser(c *gin.Context) {

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	_, err = DBconn.Exec(
		c,
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		user.Name,
		user.Email,
		string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar usuário"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário cadastrado com sucesso!"})
}

func GetUserById(c *gin.Context) {
	id := c.Param(id)

	err := DBconn.QueryRow(
		c,
		"SELECT id, name, email FROM users WHERE id=$1",
		id,
	).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param(id)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	commandTag, err := DBconn.Exec(
		c,
		"UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4",
		user.Name,
		user.Email,
		string(hashedPassword),
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}
	if commandTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso!"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param(id)

	commandTag, err := DBconn.Exec(
		c,
		"DELETE FROM users WHERE id=$1",
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
		return
	}
	if commandTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso!"})
}

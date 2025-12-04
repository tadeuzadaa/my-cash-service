package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"my-cash-service/internal/core"
	"net/http"
)

var (
	req core.LoginReq
)

func LoginUser(c *gin.Context) {

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := DBconn.QueryRow(
		c,
		"SELECT id, email, password FROM users WHERE email=$1",
		req.Email,
	).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usuário ou senha inválidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login efetuado com sucesso",
		"userId":  user.ID,
	})
}

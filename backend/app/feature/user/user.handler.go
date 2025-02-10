package featureUser

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HandleUserBalance(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// todo GetUserBalance должен принимать context первым аргументом
	// todo GetUserBalance должен принимать DTO
	user, err := GetUserBalance(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance": user.Balance,
	})
}

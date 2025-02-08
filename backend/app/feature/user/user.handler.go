package featureUser

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleUserBalance(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		c.JSON(500, gin.H{"messege": "cant convert userId"})
		return
	}

	user, err := GetUserBalance(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(500, gin.H{})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance": user.Balance,
	})
}

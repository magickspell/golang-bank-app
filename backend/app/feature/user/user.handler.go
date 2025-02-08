package featureUser

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserBalance(c *gin.Context) {
	user, err := GetUserBalance(c.Query("userId"))

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

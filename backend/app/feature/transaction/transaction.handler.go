package featureTransaction

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserTransactions(c *gin.Context) {
	transactions, err := GetTransactions(c.Query("userId"))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(500, gin.H{})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func HandleCreateTransaction(c *gin.Context) {
	err := CreateTransaction(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

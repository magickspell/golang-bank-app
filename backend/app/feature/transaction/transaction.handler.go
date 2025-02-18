package featureTransaction

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserTransactions(dbConn *sql.DB, c *gin.Context) {
	transactions, err := GetTransactions(c.Query("userId"))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"messege": "user not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func HandleCreateTransaction(dbConn *sql.DB, c *gin.Context) {
	err := CreateTransaction(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messege": "transaction complete"})
}

package featureTransaction

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserTransactions(c *gin.Context) {
	transactions, err := GetTransactions(c.Query("userId"))
	fmt.Println("transactions")
	fmt.Println(transactions)
	fmt.Println(err)

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
	CreateTransaction(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction",
	})
}

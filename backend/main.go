package main

import (
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v5/stdlib"

	transactionFeature "backend/app/feature/transaction"
	userFeature "backend/app/feature/user"
	db "backend/database"
)

// todo ENV
// todo сделать единым проброс ошибок
func main() {
	dbConn := db.Conn()
	defer dbConn.Close()

	if err := db.RunMigrations(dbConn); err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}

	// run gin app
	router := gin.Default()
	router.GET("/user-balance", userFeature.HandleUserBalance)
	router.GET("/transactions", transactionFeature.HandleUserTransactions)
	router.POST("/transactions", transactionFeature.HandleCreateTransaction)
	router.Run("localhost:8080")
}

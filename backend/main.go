package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v5/stdlib"

	transactionFeature "backend/app/feature/transaction"
	userFeature "backend/app/feature/user"
	db "backend/database"
)

// todo add graceful shutdown
// todo add logger
// todo add configLoader
// todo add auth
// todo разложить по папочкам красиво все
func main() {
	host := os.Getenv("GO_HOST")

	dbConn := db.Conn()
	defer dbConn.Close()

	if err := db.RunMigrations(dbConn); err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	} else {
		dbConn.Close()
	}

	// run gin app
	router := gin.Default()
	router.GET("/user-balance", userFeature.HandleUserBalance)
	router.GET("/transactions", transactionFeature.HandleUserTransactions)
	router.POST("/transactions", transactionFeature.HandleCreateTransaction)
	router.Run(host)
}

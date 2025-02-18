package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v5/stdlib"

	transactionFeature "backend/app/feature/transaction"
	userFeature "backend/app/feature/user"
	cfg "backend/config"
	cntx "backend/context"
	db "backend/database"
	logg "backend/logger"
)

// todo add logger
// todo add configLoader
// todo подрубить контекст
// todo сделать DTO
// todo разложить по папочкам красиво все
// todo add graceful shutdown
// todo add auth
// подрубить ормку GORM
// todo сделать нормальные тесты
func main() {
	logger := logg.NewLogger()
	config := cfg.GetConfig(logger)
	fmt.Println("Type of logger:", reflect.TypeOf(logger))
	fmt.Println("Type of config:", reflect.TypeOf(config))
	logger.OuteputLog(logg.LogPayload{Error: fmt.Errorf("error")})

	dbConn := db.Conn(config)
	defer dbConn.Close()

	if err := db.RunMigrations(dbConn); err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}
	// else {
	// 	dbConn.Close()
	// }

	// run gin app
	router := gin.Default()
	// router.Use(func(gc *gin.Context) {
	// 	gc.Set("config", config)
	// 	gc.Set("logger", logger)
	// 	gc.Next()
	// })
	router.Use(cntx.ContextMiddleware(config, logger))
	router.GET("/user-balance", func(gc *gin.Context) { userFeature.HandleUserBalance(logger, config, gc) })
	router.GET("/transactions", func(gc *gin.Context) { transactionFeature.HandleUserTransactions(dbConn, gc) })
	router.POST("/transactions", func(gc *gin.Context) { transactionFeature.HandleCreateTransaction(dbConn, gc) })
	router.Run(config.Host)
}

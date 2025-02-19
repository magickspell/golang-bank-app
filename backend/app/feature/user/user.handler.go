package featureUser

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	cfg "backend/config"
	cntx "backend/context"

	// cntx "backend/context"
	logg "backend/logger"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HandleUserBalance(logger *logg.Logger, config *cfg.Config, gc *gin.Context) {
	userId, err := strconv.Atoi(gc.Query("userId"))
	if err != nil {
		gc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx, ex := gc.Get("ctx")
	if ex {
		if ctx, ok := ctx.(*cntx.Context); ok {
			// Теперь ctx имеет тип *Context, и вы можете работать с ним
			fmt.Println(ctx.Timeout)
		} else {
			// Приведение типа не удалось, unknown не является *Context
			logger.OuteputLog(logg.LogPayload{Error: fmt.Errorf("cant type context")})
		}
		fmt.Println("[ctx]")
		fmt.Println(ctx)
		fmt.Println(&ctx)
	}
	/*
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done() // Убедимся, что wg.Done() будет вызван по завершении горутины
			time.Sleep(time.Second * 4)
		}()
		wg.Wait() // Ожидаем завершения горутины
	*/
	// Убираем горутину и WaitGroup, выполняем задержку синхронно
	time.Sleep(time.Second * 7)

	// todo GetUserBalance должен принимать context первым аргументом
	// todo GetUserBalance должен принимать DTO
	user, err := GetUserBalance(logger, config, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			gc.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			gc.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	gc.JSON(http.StatusOK, gin.H{
		"balance": user.Balance,
	})
}

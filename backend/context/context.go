package context

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v5/stdlib"

	cfg "backend/config"
	logg "backend/logger"
)

// как сделать чтобы тип gin.Context содержал	 мой конекст без приведения типа?
type Context struct {
	*gin.Context
	Config      *cfg.Config
	Logger      *logg.Logger
	Timeout     int64
	IsCancelled bool
}

func ContextMiddleware(config *cfg.Config, logger *logg.Logger) gin.HandlerFunc {
	return func(gc *gin.Context) {
		// Создаем кастомный контекст
		appContext := &Context{
			Context:     gc, // передаем в конект контекс гина
			Config:      config,
			Logger:      logger,
			Timeout:     2000, // 3 секунды
			IsCancelled: false,
		}

		// Заменяем оригинальный gin.Context на наш кастомный контекст
		*gc = *appContext.Context
		// Добавляем кастомный контекст в gin.Context
		// gc.Set("ctx", context)

		// create context with timeout
		workTime := time.Millisecond * time.Duration(appContext.Timeout)
		//timeoutCtx, cancel := context.WithTimeout(gc.Request.Context(), time.Duration(appContext.Timeout)*time.Millisecond)
		timeoutCtx, cancel := context.WithTimeout(gc.Request.Context(), workTime)
		defer cancel()
		gc.Request = gc.Request.WithContext(timeoutCtx) // change context

		// запускаем горутину с запросом
		done := make(chan int8)
		go func() {
			defer close(done)
			gc.Next()
		}()

		// Запускаем таймер для таймаута
		select {
		case <-done:
			/* Запрос завершен успешно */
		case <-timeoutCtx.Done():
			err := fmt.Errorf("request timeout")
			appContext.Logger.OuteputLog(logg.LogPayload{Error: err})
			gc.AbortWithStatusJSON(408, gin.H{"error": err.Error()})
		case <-gc.Done():
			fmt.Println("[IsCancelled]")
			// Запрос был завершен или отменен
			appContext.IsCancelled = true
			fmt.Println(appContext.IsCancelled)
			fmt.Println("[IsCancelled]")
			gc.AbortWithStatusJSON(408, gin.H{"error": "errCanceld"})
		}
	}
}

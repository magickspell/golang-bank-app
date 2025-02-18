package context

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v5/stdlib"

	cfg "backend/config"
	logg "backend/logger"
)

// как сделать чтобы тип gin.Context содержал мой конекст без приведения типа?
type Context struct {
	Config      *cfg.Config
	Logger      *logg.Logger
	Timeout     int64
	IsCancelled bool
}

type GinContext struct {
	GinCtx *gin.Context
	AppCtx Context
}

func ContextMiddleware(config *cfg.Config, logger *logg.Logger) gin.HandlerFunc {
	return func(gc *gin.Context) {
		// Создаем кастомный контекст
		context := &Context{
			Config:      config,
			Logger:      logger,
			Timeout:     3000, // 3 секунды
			IsCancelled: false,
		}

		// Добавляем кастомный контекст в gin.Context
		gc.Set("ctx", context)

		// Запускаем таймер для таймаута
		go func() {
			select {
			case <-time.After(time.Duration(context.Timeout) * time.Millisecond):
				if !context.IsCancelled {
					err := fmt.Errorf("request timeout")
					context.Logger.OuteputLog(logg.LogPayload{Error: err})
					gc.AbortWithStatusJSON(408, gin.H{"error": err.Error()})
				}
			case <-gc.Done():
				// Запрос был завершен или отменен
				context.IsCancelled = true
			}
		}()

		gc.Next()
	}
}

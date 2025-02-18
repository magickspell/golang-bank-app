package featureUser

import (
	"database/sql"
	"net/http"
	"strconv"

	cfg "backend/config"
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

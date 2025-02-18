package featureUser

import (
	cfg "backend/config"
	logg "backend/logger"
)

func GetUserBalance(logger *logg.Logger, config *cfg.Config, userId int) (User, error) {
	// todo GetUser должен принимать context первым аргументом (context.Context)
	return GetUser(logger, config, userId)
}

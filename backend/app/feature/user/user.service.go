package featureUser

func GetUserBalance(userId int) (User, error) {
	// todo GetUser должен принимать context первым аргументом (context.Context)
	return GetUser(userId)
}

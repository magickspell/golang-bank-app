package featureUser

func GetUserBalance(userId int) (User, error) {
	return GetUser(userId)
}

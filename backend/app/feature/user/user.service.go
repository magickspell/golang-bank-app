package featureUser

func GetUserBalance(userId string) (User, error) {
	return GetUser(userId)
}

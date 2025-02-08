package featureUser

import (
	db "backend/database"
	"fmt"
	"os"
)

type User struct {
	Id      int
	Balance int
}

type Transaction struct {
	Id        int
	Amount    int
	ToUser    int
	FromUser  *int
	CreatedAt string
}

func GetUser(userId string) (User, error) {
	dbConn := db.Conn()
	defer dbConn.Close()

	var id, balance int

	err := dbConn.QueryRow("SELECT * FROM users u WHERE u.id=$1;", userId).Scan(&id, &balance)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return User{}, err
	}

	return User{
		Id:      id,
		Balance: balance,
	}, nil
}

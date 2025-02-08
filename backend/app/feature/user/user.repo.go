package featureUser

import (
	db "backend/database"
	"database/sql"
	"fmt"
	"os"
)

type User struct {
	Id      int
	Balance int
}

type Operation string

const (
	OPERATION_PLUS  Operation = "plus"
	OPERATION_MINUS Operation = "minus"
)

func GetUser(userId int) (User, error) {
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

func UpdateUserBalance(userId int, amount int, operation Operation, tran *sql.Tx) error {
	dbConn := db.Conn()
	defer dbConn.Close()

	var query string
	if operation == OPERATION_PLUS {
		query = "UPDATE users SET balance = balance + $1 WHERE id = $2"
	}
	if operation == OPERATION_MINUS {
		query = "UPDATE users SET balance = balance - $1 WHERE id = $2"
	}

	_, err := tran.Exec(query, amount, userId)
	if err != nil {
		fmt.Println("cant process transaction")
		fmt.Println(err)
		return err
	}

	return nil
}

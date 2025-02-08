package featureTransaction

import (
	db "backend/database"
	"fmt"
)

type Transaction struct {
	Id        int
	Amount    int
	ToUser    int
	FromUser  *int
	CreatedAt string
}

func GetUserTransactions(userId string) ([]Transaction, error) {
	dbConn := db.Conn()
	defer dbConn.Close()

	rows, err := dbConn.Query(
		"SELECT * FROM transactions t WHERE t.user_to=$1 OR t.user_from=$1;",
		userId,
	)
	if err != nil {
		fmt.Println("ошибка при сканировании строки: %v", err)
		return []Transaction{}, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.Id, &transaction.ToUser, &transaction.FromUser, &transaction.Amount, &transaction.CreatedAt)
		if err != nil {
			fmt.Println("ошибка при сканировании строки: %v", err)
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}
		fmt.Println(transaction)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func InsertTransaction(amount int, userFrom string, userTo string) error {
	dbConn := db.Conn()
	defer dbConn.Close()

	tran, err := dbConn.Begin()
	if err != nil {
		fmt.Println("cant start transaction")
		return err
	}

	// TODO
	// минус баланс у FROM
	// плюс баланс у TO
	// записали транзакцию
	_, err = tran.Exec(
		"INSERT INTO transactions (amount, user_from, user_to) VALUES ($1, $2, $3);",
		amount, userFrom, userTo,
	)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

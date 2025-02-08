package featureTransaction

import (
	user "backend/app/feature/user"
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
		`
		SELECT * FROM transactions t
		WHERE t.user_to=$1
			OR t.user_from=$1
		ORDER BY created_at desc
		LIMIT 10
		;
		`,
		userId,
	)
	if err != nil {
		fmt.Printf("ошибка при сканировании строки: %v", err)
		return []Transaction{}, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.Id, &transaction.ToUser, &transaction.FromUser, &transaction.Amount, &transaction.CreatedAt)
		if err != nil {
			fmt.Printf("ошибка при сканировании строки: %v", err)
			return nil, fmt.Errorf("ошибка при сканировании строки: %v", err)
		}
		fmt.Println(transaction)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func InsertTransaction(amount int, userFrom *int, userTo int) error {
	dbConn := db.Conn()
	defer dbConn.Close()

	// стартуем транзакцию по переводу денег от одного пользователя к другому
	tran, err := dbConn.Begin()
	if err != nil {
		fmt.Println("cant start transaction")
		return err
	}

	if userFrom != nil {
		err = user.UpdateUserBalance(*userFrom, amount, user.OPERATION_MINUS, tran)
		if err != nil {
			fmt.Println("cant OPERATION_MINUS")
			return err
		}
	}
	err = user.UpdateUserBalance(userTo, amount, user.OPERATION_PLUS, tran)
	if err != nil {
		fmt.Println("cant OPERATION_PLUS")
		return err
	}

	_, err = tran.Exec(
		"INSERT INTO transactions (amount, user_from, user_to) VALUES ($1, $2, $3);",
		amount, userFrom, userTo,
	)
	if err != nil {
		fmt.Println("unable to insert row: %w", err)
		return fmt.Errorf("unable to insert row: %w", err)
	}

	// комитим транзакцию
	err = tran.Commit()
	if err != nil {
		fmt.Println("unable to commit row: %w", err)
		return fmt.Errorf("unable to commit row: %w", err)
	}

	return nil
}

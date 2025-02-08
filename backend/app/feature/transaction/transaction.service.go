package featureTransaction

import (
	"fmt"

	user "backend/app/feature/user"

	"github.com/gin-gonic/gin"
)

type transactionRequest struct {
	Amount     int  `json:"amount"`
	UserFromId *int `json:"userFromId"`
	UserToId   int  `json:"userToId"`
}

func GetTransactions(userId string) ([]Transaction, error) {
	return GetUserTransactions(userId)
}

func CreateTransaction(c *gin.Context) error {
	var req transactionRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("error: cant process transaction")
		fmt.Println(err)
		return fmt.Errorf("error: cant process transaction")
	}

	if req.UserToId == 0 || req.Amount == 0 {
		fmt.Println("Error: some properties wasnot provided")
		return fmt.Errorf("error: some properties wasnot provided")
	}

	_, err := user.GetUserBalance(req.UserToId)
	if err != nil {
		fmt.Println("Cant find user [UserToId]:", err)
		return err
	}

	if req.UserFromId != nil {
		userFrom, err := user.GetUserBalance(*req.UserFromId)
		if err != nil {
			fmt.Println("Cant find user [UserFromId]:", err)
			return err
		}
		if userFrom.Balance < req.Amount {
			fmt.Println("баланс меньше суммы перевода")
			return fmt.Errorf("баланс меньше суммы перевода")
		}
	}

	InsertTransaction(req.Amount, req.UserFromId, req.UserToId)
	return nil
}

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
		return fmt.Errorf("error: cant process transaction request")
	}

	if req.UserToId == 0 || req.Amount == 0 {
		return fmt.Errorf("error: some properties wasnot provided")
	}

	_, err := user.GetUserBalance(req.UserToId)
	if err != nil {
		return fmt.Errorf("cant find user [UserToId]: '%v'", err)
	}

	if req.UserFromId != nil {
		userFrom, err := user.GetUserBalance(*req.UserFromId)
		if err != nil {
			return fmt.Errorf("cant find user [UserFromId]: '%v'", err)
		}
		if userFrom.Balance < req.Amount {
			return fmt.Errorf("баланс меньше суммы перевода")
		}
	}

	err = InsertTransaction(req.Amount, req.UserFromId, req.UserToId)
	if err != nil {
		return fmt.Errorf("cant InsertTransaction: '%v'", err)
	}
	return nil
}

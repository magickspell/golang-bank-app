package featureTransaction

import (
	"fmt"
	"strconv"

	user "backend/app/feature/user"

	"github.com/gin-gonic/gin"
)

func GetTransactions(userId string) ([]Transaction, error) {
	return GetUserTransactions(userId)
}

func CreateTransaction(c *gin.Context) error {
	userFromId := c.PostForm("userFromId")
	userToId := c.PostForm("userToId")

	strAmount := c.PostForm("amount")
	fmt.Println("potFORM")
	fmt.Println(userFromId)
	fmt.Println(userToId)
	fmt.Println(strAmount)
	amount, err := strconv.Atoi(strAmount)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	userFrom, err := user.GetUserBalance(userFromId)
	if err != nil {
		fmt.Println("Cant find user:", err)
		return err
	}
	fmt.Println(userFrom)
	if userFrom.Balance < amount {
		fmt.Println("Баланс меньше суммы перевода")
		return fmt.Errorf("Баланс меньше суммы перевода")
	}

	InsertTransaction(amount, userFromId, userToId)
	return nil
}

package services

import (
	. "internal-account-bank/models"
	"time"
)

func GetAllAccounts() ([]Account, error) {
	accounts, err := dao.GetAllAccounts()
	return accounts, err
}

func GetAccountById(id string) (Account, error) {
	account, err := dao.GetAccountById(id)

	return account, err
}

func GetBalanceByAccountID(id string) (float32, error) {
	balance, err := dao.GetBalanceByAccountID(id)

	return balance, err
}

func CreateAccount(account Account) error {
	account.Created_At = time.Now()
	return dao.CreateAccount(account)
}
package services

import (
	. "internal-account-bank/dao"
	. "internal-account-bank/models"
	"time"
)

var dao = AccountsDAO{}

func init() {
	config := dbConfig()
	dao.Server = config[dbUrl]
	dao.Database = config[dbName]
	dao.Connect()
}

func GetAllAccounts() ([]Account, error) {
	accounts, err := dao.GetAllAccounts()
	return accounts, err
}

func GetBalanceByAccountID(id string) (float32, error) {
	balance, err := dao.GetBalanceByAccountID(id)

	return balance, err
}

func CreateAccount(account Account) error {
	account.Created_At = time.Now()
	return dao.CreateAccount(account)
}

package services

import (
	"errors"
	. "internal-account-bank/models"

	"gopkg.in/mgo.v2/bson"
)

func GetAllAccounts() ([]Account, error) {
	accounts, err := dao.GetAllAccounts()
	return accounts, err
}

func GetAccountById(id string) (Account, error) {
	account, err := dao.GetAccountById(id)

	return account, err
}

func GetAccountByCpf(cpf string) (Account, error) {
	account, err := dao.GetAccountByCpf(cpf)

	return account, err
}

func GetBalanceByAccountID(id string) (float32, error) {
	if !bson.IsObjectIdHex(id) {
		return 0, errors.New("Account ID Invalid.")
	}
	balance, err := dao.GetBalanceByAccountID(id)

	return balance, err
}

func CreateAccount(account Account) error {
	foundAccount, err := GetAccountByCpf(account.Cpf)
	if foundAccount.Cpf != "" || err != nil && err.Error() != "not found" {
		return errors.New("Account with this cpf already exists.")
	}
	return dao.CreateAccount(account)
}

func UpdateAccount(id bson.ObjectId, account Account) error {
	return dao.UpdateAccount(id, account)
}

package dao

import (
	. "internal-account-bank/models"

	"gopkg.in/mgo.v2/bson"
)

func (m *DAO) GetAllAccounts() ([]Account, error) {
	var account []Account
	err := db.C(COLLECTION_ACCOUNTS).Find(bson.M{}).All(&account)
	return account, err
}

func (m *DAO) GetAccountById(id string) (Account, error) {
	var account Account
	err := db.C(COLLECTION_ACCOUNTS).FindId(bson.ObjectIdHex(id)).One(&account)
	return account, err
}

func (m *DAO) GetAccountByCpf(cpf string) (Account, error) {
	var account Account
	err := db.C(COLLECTION_ACCOUNTS).Find(bson.M{"cpf": cpf}).Select(bson.M{"cpf": 1}).One(&account)
	return account, err
}

func (m *DAO) GetBalanceByAccountID(id string) (float32, error) {
	var account Account
	err := db.C(COLLECTION_ACCOUNTS).FindId(bson.ObjectIdHex(id)).Select(bson.M{"balance": 1}).One(&account)
	return account.Balance, err
}

func (m *DAO) CreateAccount(account Account) error {
	err := db.C(COLLECTION_ACCOUNTS).Insert(&account)
	return err
}

func (m *DAO) UpdateAccount(id bson.ObjectId, account Account) error {
	err := db.C(COLLECTION_ACCOUNTS).UpdateId(id, &account)
	return err
}

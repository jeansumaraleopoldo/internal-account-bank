package dao

import (
	"log"

	. "internal-account-bank/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountsDAO struct {
	Server   string
	Database string
}

func (m *AccountsDAO) Connect() {
	log.Println(m.Server)
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *AccountsDAO) GetAllAccounts() ([]Account, error) {
	var account []Account
	err := db.C(COLLECTION_ACCOUNTS).Find(bson.M{}).All(&account)
	return account, err
}

func (m *AccountsDAO) GetBalanceByAccountID(id string) (float32, error) {
	var account Account
	err := db.C(COLLECTION_ACCOUNTS).FindId(bson.ObjectIdHex(id)).Select(bson.M{"balance": 1}).One(&account)
	return account.Balance, err
}

func (m *AccountsDAO) CreateAccount(account Account) error {
	err := db.C(COLLECTION_ACCOUNTS).Insert(&account)
	return err
}

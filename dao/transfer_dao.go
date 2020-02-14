package dao

import (
	"log"

	. "github.com/jeansumaraleopoldo/internal-account-bank/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TransfersDAO struct {
	Server   string
	Database string
}

func (m *TransfersDAO) Connect() {
	log.Println(m.Server)
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TransfersDAO) GetAllTransfers() ([]Account, error) {
	var account []Account
	err := db.C(COLLECTION_TRANSFERS).Find(bson.M{}).All(&account)
	return account, err
}

func (m *AccountsDAO) CreateTransfer(account Account) error {
	err := db.C(COLLECTION_TRANSFERS).Insert(&account)
	return err
}

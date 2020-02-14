package dao

import (
	"log"

	. "github.com/jeansumaraleopoldo/internal-account-bank/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "accounts"
)

func (m *AccountsDAO) Connect() {
	log.Println(m.Server)
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *AccountsDAO) GetAll() ([]Account, error) {
	var account []Account
	err := db.C(COLLECTION).Find(bson.M{}).All(&account)
	return account, err
}

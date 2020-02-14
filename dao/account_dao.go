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

func (m *AccountsDAO) GetByID(id string) (Account, error) {
	var account Account
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&account)
	return account, err
}

func (m *AccountsDAO) Create(account Account) error {
	err := db.C(COLLECTION).Insert(&account)
	return err
}

func (m *AccountsDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *AccountsDAO) Update(id string, account Account) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &account)
	return err
}

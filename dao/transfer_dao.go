package dao

import (
	. "internal-account-bank/models"

	"gopkg.in/mgo.v2/bson"
)

func (m *DAO) GetAllTransfers() ([]Transfer, error) {
	var transfer []Transfer
	err := db.C(COLLECTION_TRANSFERS).Find(bson.M{}).All(&transfer)
	return transfer, err
}

func (m *DAO) CreateTransfer(transfer Transfer) error {
	err := db.C(COLLECTION_TRANSFERS).Insert(&transfer)
	return err
}

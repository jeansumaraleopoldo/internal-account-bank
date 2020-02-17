package services

import (
	. "internal-account-bank/dao"
)

var dao = DAO{}

func init() {
	dao.Server = "mongo:27017"
	dao.Database = "bank"
	dao.Connect()
}

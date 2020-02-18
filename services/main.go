package services

import (
	. "internal-account-bank/dao"
)

var dao = DAO{}

func init() {
	//dao.Server = "mongo:27017"
	dao.Server = "172.18.0.3:27017"
	dao.Database = "bank"
	dao.Connect()
}

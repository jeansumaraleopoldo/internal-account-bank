package controller

import (
	. "github.com/jeansumaraleopoldo/internal-account-bank/dao"
	. "github.com/jeansumaraleopoldo/internal-account-bank/models"
)

var dao = AccountsDAO{}

const (
	dbUrl  = "DB_URL"
	dbName = "DB_NAME"
)

func init() {
	config := dbConfig()
	dao.Server = config[dbUrl]
	dao.Database = config[dbName]
	dao.Connect()
}

func dbConfig() map[string]string {
	conf := make(map[string]string)

	conf[dbUrl] = "localhost"
	conf[dbName] = "bank"
	return conf
}

func getAll() ([]Account, error) {
	accounts, err := dao.GetAll()

	return accounts, err
}

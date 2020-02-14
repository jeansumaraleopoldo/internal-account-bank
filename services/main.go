package services

import (
	. "internal-account-bank/dao"
)

const (
	dbUrl  = "DB_URL"
	dbName = "DB_NAME"
)

var dao = DAO{}

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

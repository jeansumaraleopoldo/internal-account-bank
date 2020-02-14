package services

import (
	. "internal-account-bank/dao"
	. "internal-account-bank/models"
	"time"
)

var daoTransfer = TransfersDAO{}

func init() {
	config := dbConfig()
	daoTransfer.Server = config[dbUrl]
	daoTransfer.Database = config[dbName]
	daoTransfer.Connect()
}

func GetAllTransfers() ([]Transfer, error) {
	accounts, err := daoTransfer.GetAllTransfers()
	return accounts, err
}

func CreateTransfer(transfer Transfer) error {
	transfer.Created_At = time.Now()
	return daoTransfer.CreateTransfer(transfer)
}

func TransferBetweenAccount(transfer Transfer) {

}

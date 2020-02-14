package services

import (
	"errors"
	. "internal-account-bank/models"
	"time"
)

func GetAllTransfers() ([]Transfer, error) {
	accounts, err := dao.GetAllTransfers()
	return accounts, err
}

func CreateTransfer(transfer Transfer) error {
	transfer.Created_At = time.Now()
	return dao.CreateTransfer(transfer)
}

func TransferBetweenAccount(transfer Transfer) error {
	if transfer.Account_Origin_Id == transfer.Account_Destination_Id {
		return errors.New("Nao é possível fazer uma transferencia para a mesma conta.")
	}

	accountOrigin, err := GetAccountById(transfer.Account_Origin_Id)

	if accountOrigin.Balance == 0 || transfer.Amount > accountOrigin.Balance {
		return errors.New("O saldo da conta de origem é insuficiente para fazer uma transferencia.")
	}

	accountDestination, err := GetAccountById(transfer.Account_Destination_Id)

	accountOrigin.Balance = accountOrigin.Balance - transfer.Amount
	accountDestination.Balance = accountDestination.Balance + transfer.Amount

	if err := UpdateAccount(accountOrigin.ID, accountOrigin); err != nil {
		return err
	}

	if err := UpdateAccount(accountDestination.ID, accountDestination); err != nil {
		return err
	}

	return err
}

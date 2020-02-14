package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	. "internal-account-bank/models"

	"gopkg.in/mgo.v2/bson"

	"internal-account-bank/services"
)

func TransferGetAll(w http.ResponseWriter, r *http.Request) {
	transfers, err := services.GetAllTransfers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, transfers)
}

func TransferCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var transfer Transfer
	transfer.Created_At = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	transfer.ID = bson.NewObjectId()

	if err := services.TransferBetweenAccount(transfer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := services.CreateTransfer(transfer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, transfer)
}

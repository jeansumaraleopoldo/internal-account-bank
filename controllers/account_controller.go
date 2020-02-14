package controller

import (
	"encoding/json"
	"net/http"

	. "github.com/jeansumaraleopoldo/internal-account-bank/dao"
	. "github.com/jeansumaraleopoldo/internal-account-bank/models"
)

var dao = AccountsDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AccountGetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, accounts)
}

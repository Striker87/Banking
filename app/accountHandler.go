package app

import (
	"encoding/json"
	"net/http"

	"github.com/Striker87/Banking/dto"
	"github.com/Striker87/Banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = mux.Vars(r)["customer_id"]
	account, appError := h.service.NewAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}

	writeResponse(w, http.StatusCreated, account)
}

package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Striker87/Banking/service"
	"github.com/gorilla/mux"
)

type CustomerHanlders struct {
	service service.CustomerService
}

func (ch *CustomerHanlders) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHanlders) getCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["customer_id"]
	w.Header().Set("Content-Type", "application/json")

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("writeResponse() failed due error: %v", err)
	}
}

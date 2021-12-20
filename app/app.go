package app

import (
	"log"
	"net/http"

	"github.com/Striker87/Banking/domain"
	"github.com/Striker87/Banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func Start() {
	router := mux.NewRouter()

	//ch := CustomerHanlders{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHanlders{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

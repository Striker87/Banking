package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Striker87/Banking/domain"
	"github.com/Striker87/Banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDR") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable SERVER_ADDR or SERVER_PORT not defined")
	}
}

func Start() {
	//sanityCheck()
	router := mux.NewRouter()

	//ch := CustomerHanlders{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	db := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(db)
	accountRepositoryDb := domain.NewAccountRepositoryDb(db)

	ch := CustomerHanlders{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	//address := os.Getenv("SERVER_ADDRESS")
	//port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", "", "8080"), router))
}

func getDbClient() *sqlx.DB {
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbAddr := os.Getenv("DB_ADDR")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "", "127.0.0.1", "3306", "banking")
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

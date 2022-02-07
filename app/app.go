package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/poinguardvinay/goweb/domain"
	"github.com/poinguardvinay/goweb/service"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()
	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//router.HandleFunc("/cat", catHandler).Methods(http.MethodGet)

	//router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	//router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8989", router))
}
